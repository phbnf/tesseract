## 2026-06-16 - Prevent Profiling Endpoint Exposure in Public Entry Points

**Vulnerability:** Anonymous imports of `net/http/pprof` (in `cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go`) and `expvar` (in `cmd/tesseract/posix/main.go`) automatically expose profiling and metrics endpoints (`/debug/pprof/*` and `/debug/vars`) on the global `http.DefaultServeMux`. If `http.DefaultServeMux` is used or accessible publicly, this exposes internal application details, leading to a CWE-200 (Information Exposure) vulnerability.

**Learning:** When creating cloud personalities or public-facing HTTP servers, it is crucial to avoid anonymous imports that mutate the default multiplexer. These endpoints provide valuable insight into the application's runtime state, which can be leveraged by attackers for reconnaissance or denial-of-service (DoS) attacks.

**Prevention:** Never use anonymous imports like `_ "net/http/pprof"` or `_ "expvar"` in entry points that start HTTP servers for public API access. If profiling is strictly required, it should be bound to a separate, internal administrative server or port, not the public-facing listener.
