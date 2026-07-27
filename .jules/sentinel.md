## 2025-03-25 - [HIGH] Fix CWE-200 Profiling Endpoint Exposure
**Vulnerability:** Profiling endpoint (`/debug/pprof`) is automatically exposed on public HTTP servers because the `net/http/pprof` package is imported in `cmd/tesseract/gcp/main.go` and `cmd/tesseract/posix/main.go`.
**Learning:** Importing `net/http/pprof` for side effects registers its handlers to `http.DefaultServeMux`. If `http.DefaultServeMux` is inadvertently used or if the public server mux isn't strictly isolated from it, it exposes sensitive process internal state (CWE-200).
**Prevention:** Do not import `_ "net/http/pprof"` in production binaries that expose public HTTP endpoints unless it is explicitly bound to a private, non-public listener.
