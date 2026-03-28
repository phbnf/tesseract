## 2024-03-28 - [Profiler Exposed]
**Vulnerability:** Automatically exposed `net/http/pprof` endpoints.
**Learning:** `_ "net/http/pprof"` was anonymously imported in `cmd/tesseract/posix/main.go`, inadvertently exposing profiling endpoints via `http.DefaultServeMux`, which introduces a CWE-200 vulnerability when the mux is exposed.
**Prevention:** Remove `_ "net/http/pprof"` imports from public-facing binaries and configure pprof to use a dedicated, internal HTTP server if profiling is required.
