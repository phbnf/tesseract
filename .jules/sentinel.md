## 2025-05-14 - Prevent unintentional exposure of profiling data
**Vulnerability:** Anonymous imports of `net/http/pprof` and `expvar` in `cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go` exposed profiling data (`/debug/pprof/*` and `/debug/vars`) on the default HTTP multiplexer (CWE-200).
**Learning:** Using `_ "net/http/pprof"` and `_ "expvar"` automatically registers sensitive metrics endpoints globally, making them accessible to any client hitting the service, unless the server explicitly defines a custom routing mux that excludes them.
**Prevention:** Never use anonymous imports for pprof or expvar in production entry points unless explicitly configuring them on a secure, internal-only HTTP server.
