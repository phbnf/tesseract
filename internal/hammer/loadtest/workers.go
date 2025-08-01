// Copyright 2024 The Tessera authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loadtest

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/transparency-dev/formats/log"
	"github.com/transparency-dev/merkle/proof"
	hasher "github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tesseract/internal/client"
	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
	"github.com/transparency-dev/tesseract/internal/x509util"
	"k8s.io/klog/v2"
)

// LeafWriter is the signature of a function which can write arbitrary data to a log.
// The data to be written is provided, and the implementation must return the sequence
// number at which this data will be found in the log and the timestamp of the SCT
// issued for the data, or an error.
type LeafWriter func(ctx context.Context, data []byte) (index uint64, timestamp uint64, err error)

type LogReader interface {
	ReadCheckpoint(ctx context.Context) ([]byte, error)

	ReadTile(ctx context.Context, l, i uint64, p uint8) ([]byte, error)

	ReadEntryBundle(ctx context.Context, i uint64, p uint8) ([]byte, error)
}

// NewLeafReader creates a LeafReader.
// The next function provides a strategy for which leaves will be read.
// Custom implementations can be passed, or use RandomNextLeaf or MonotonicallyIncreasingNextLeaf.
func NewLeafReader(tracker *client.LogStateTracker, f client.EntryBundleFetcherFunc, next func(uint64) uint64, throttle <-chan bool, errChan chan<- error) *LeafReader {
	return &LeafReader{
		tracker:  tracker,
		f:        f,
		next:     next,
		throttle: throttle,
		errChan:  errChan,
	}
}

// LeafReader reads leaves from the tree.
// This class is not thread safe.
type LeafReader struct {
	tracker  *client.LogStateTracker
	f        client.EntryBundleFetcherFunc
	next     func(uint64) uint64
	throttle <-chan bool
	errChan  chan<- error
	cancel   func()
	c        leafBundleCache
}

// Run runs the log reader. This should be called in a goroutine.
func (r *LeafReader) Run(ctx context.Context) {
	if r.cancel != nil {
		panic("LeafReader was ran multiple times")
	}
	ctx, r.cancel = context.WithCancel(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case <-r.throttle:
		}
		size := r.tracker.LatestConsistent.Size
		if size == 0 {
			continue
		}
		i := r.next(size)
		if i >= size {
			continue
		}
		klog.V(2).Infof("LeafReader getting %d", i)
		_, err := r.getLeaf(ctx, i, size)
		if err != nil {
			r.errChan <- fmt.Errorf("failed to get leaf %d: %v", i, err)
		}
	}
}

// getLeaf fetches the raw contents committed to at a given leaf index.
func (r *LeafReader) getLeaf(ctx context.Context, i uint64, logSize uint64) ([]byte, error) {
	if i >= logSize {
		return nil, fmt.Errorf("requested leaf %d >= log size %d", i, logSize)
	}
	if cached, _ := r.c.get(i); cached != nil {
		klog.V(2).Infof("Using cached result for index %d", i)
		return cached, nil
	}

	bundle, err := client.GetEntryBundle(ctx, r.f, i/layout.EntryBundleWidth, logSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get entry bundle: %v", err)
	}
	ti := i % layout.EntryBundleWidth
	r.c = leafBundleCache{
		start:  i - ti,
		leaves: bundle.Entries,
	}
	return r.c.leaves[ti], nil
}

// Kills this leaf reader at the next opportune moment.
// This function may return before the reader is dead.
func (r *LeafReader) Kill() {
	if r.cancel != nil {
		r.cancel()
	}
}

// leafBundleCache stores the results of the last fetched tile. This allows
// readers that read contiguous blocks of leaves to act more like real
// clients and fetch a tile of 256 leaves once, instead of 256 times.
type leafBundleCache struct {
	start  uint64
	leaves [][]byte
}

func (tc leafBundleCache) get(i uint64) ([]byte, error) {
	end := tc.start + uint64(len(tc.leaves))
	if i >= tc.start && i < end {
		leaf := tc.leaves[i-tc.start]
		return base64.StdEncoding.DecodeString(string(leaf))
	}
	return nil, errors.New("not found")
}

// RandomNextLeaf returns a function that fetches a random leaf available in the tree.
func RandomNextLeaf() func(uint64) uint64 {
	return func(size uint64) uint64 {
		return rand.Uint64N(size)
	}
}

// MonotonicallyIncreasingNextLeaf returns a function that always wants the next available
// leaf after the one it previously fetched. It starts at leaf 0.
func MonotonicallyIncreasingNextLeaf() func(uint64) uint64 {
	var i uint64
	return func(size uint64) uint64 {
		if i < size {
			r := i
			i++
			return r
		}
		return size
	}
}

// LeafTime records the time at which a leaf was assigned the given index.
//
// This is used when sampling leaves which are added in order to later calculate
// how long it took to for them to become integrated.
type LeafTime struct {
	Index      uint64
	QueuedAt   time.Time
	AssignedAt time.Time
}

// LeafMMD records the generated leaf in the request and the
// timestamp in the response.
//
// This is used to verify the MMD violation by the performing the inclusion proof.
type LeafMMD struct {
	leaf      []byte
	index     uint64
	timestamp uint64
}

// NewLogWriter creates a LogWriter.
// u is the URL of the write endpoint for the log.
// gen is a function that generates new leaves to add.
func NewLogWriter(writer LeafWriter, gen func() rfc6962.AddChainRequest, throttle <-chan bool, errChan chan<- error, leafSampleChan chan<- LeafTime, leafMMDChan chan<- LeafMMD) *LogWriter {
	return &LogWriter{
		writer:      writer,
		gen:         gen,
		throttle:    throttle,
		errChan:     errChan,
		leafChan:    leafSampleChan,
		leafMMDChan: leafMMDChan,
	}
}

// LogWriter writes new leaves to the log that are generated by `gen`.
type LogWriter struct {
	writer LeafWriter
	gen    func() rfc6962.AddChainRequest

	throttle    <-chan bool
	errChan     chan<- error
	leafChan    chan<- LeafTime
	leafMMDChan chan<- LeafMMD
	cancel      func()
}

// Run runs the log writer. This should be called in a goroutine.
func (w *LogWriter) Run(ctx context.Context) {
	if w.cancel != nil {
		panic("LogWriter was run multiple times")
	}
	ctx, w.cancel = context.WithCancel(ctx)
	newLeaf := w.gen()
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.throttle:
		}
		reqBody, err := json.Marshal(newLeaf)
		if err != nil {
			klog.Errorf("Failed to json.Marshal add chain request body: %v", err)
			continue
		}

		lt := LeafTime{QueuedAt: time.Now()}
		index, timestamp, err := w.writer(ctx, reqBody)
		if err != nil {
			w.errChan <- fmt.Errorf("failed to create request: %w", err)
			continue
		}
		lt.Index, lt.AssignedAt = index, time.Now()

		// See if we can send a leaf sample.
		select {
		// TODO: we might want to count dropped samples, and/or make sampling a bit more statistical.
		case w.leafChan <- lt:
		default:
		}

		var chain []byte
		for _, cert := range newLeaf.Chain {
			chain = append(chain, cert...)
		}

		// Send LeafMMD for inclusion proof verification.
		if cap(w.leafMMDChan) > 0 {
			select {
			case w.leafMMDChan <- LeafMMD{chain, index, timestamp}:
			default:
				// Drop if leafMMDChan is full. This could happen if the MMD verifiers are falling behind.
				klog.V(3).Infof("leafMMDChan is full: dropping leaf index: %d", index)
			}
		}

		klog.V(2).Infof("Wrote leaf at index %d", index)
		newLeaf = w.gen()
	}
}

// Kills this writer at the next opportune moment.
// This function may return before the writer is dead.
func (w *LogWriter) Kill() {
	if w.cancel != nil {
		w.cancel()
	}
}

// NewMMDVerifier creates a MMDVerifier.
func NewMMDVerifier(tracker *client.LogStateTracker, mmdDuration time.Duration, errChan chan<- error, leafMMDChan <-chan LeafMMD) *MMDVerifier {
	return &MMDVerifier{
		tracker:     tracker,
		mmdDuration: mmdDuration,
		errChan:     errChan,
		leafMMDChan: leafMMDChan,
	}
}

// MMDVerifier verifies the signed timestamp against the MMD policy for newly
// added entries by performing inclusion proof.
type MMDVerifier struct {
	tracker     *client.LogStateTracker
	mmdDuration time.Duration
	errChan     chan<- error
	leafMMDChan <-chan LeafMMD
	cancel      func()
}

// Run runs the MMD verifier. This should be called in a goroutine.
func (v *MMDVerifier) Run(ctx context.Context) {
	if v.cancel != nil {
		panic("MMDVerifier was ran multiple times")
	}
	ctx, v.cancel = context.WithCancel(ctx)

	var checkpoint log.Checkpoint
	var proofBuilder *client.ProofBuilder
	var leafMMD *LeafMMD // The leaf we are currently processing, if any.
	for {
		if leafMMD == nil {
			select {
			case <-ctx.Done():
				return
			case mmd, ok := <-v.leafMMDChan:
				if !ok {
					return
				}
				leafMMD = &mmd
			}
		} else {
			// We have a leaf but failed to find it integrated last time, wait a bit and try again.
			time.Sleep(100 * time.Millisecond)
		}

		// Retry if the leaf is not yet integrated into the log.
		if leafMMD.index >= v.tracker.LatestConsistent.Size {
			// Verify MMD timestamp. Retry until the leaf is integrated, so we know how long it violates the MMD.
			if time.UnixMilli(int64(leafMMD.timestamp)).Add(v.mmdDuration).Before(time.Now()) {
				v.errChan <- fmt.Errorf("leaf index %d MMD violation at %d", leafMMD.index, leafMMD.timestamp)
				continue
			}

			continue
		}

		// Update the checkpoint and proof builder only if needed, so the node
		// cache in the proof builder can be effectively used to improve the
		// performance.
		if (leafMMD.index >= checkpoint.Size && checkpoint.Size < v.tracker.LatestConsistent.Size) || proofBuilder == nil {
			checkpoint = v.tracker.LatestConsistent

			var err error
			proofBuilder, err = client.NewProofBuilder(ctx, log.Checkpoint{
				Origin: v.tracker.Origin,
				Size:   checkpoint.Size,
				Hash:   checkpoint.Hash,
			}, v.tracker.TileFetcher)
			if err != nil {
				v.errChan <- fmt.Errorf("failed to create proof builder: %v", err)
				continue
			}
		}

		ip, err := proofBuilder.InclusionProof(ctx, leafMMD.index)
		if err != nil {
			v.errChan <- fmt.Errorf("failed to create inclusion proof: %v", err)
			continue
		}
		certs, err := x509.ParseCertificates(leafMMD.leaf)
		if err != nil {
			v.errChan <- fmt.Errorf("failed to parse certificates: %v", err)
			continue
		}
		entry, err := x509util.EntryFromChain(certs, false, leafMMD.timestamp)
		if err != nil {
			v.errChan <- fmt.Errorf("failed to create entry from chain: %v", err)
			continue
		}
		leafHash := entry.MerkleLeafHash(leafMMD.index)
		if err := proof.VerifyInclusion(hasher.DefaultHasher, leafMMD.index, checkpoint.Size, leafHash, ip, checkpoint.Hash); err != nil {
			v.errChan <- fmt.Errorf("failed to verify inclusion proof: %v", err)
			continue
		}

		leafMMD = nil
	}
}

// Kills this verifier at the next opportune moment.
// This function may return before the verifier is dead.
func (v *MMDVerifier) Kill() {
	if v.cancel != nil {
		v.cancel()
	}
}
