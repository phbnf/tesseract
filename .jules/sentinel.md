## 2025-02-23 - Prevent Exposing Profiling and Metrics Endpoints

**Vulnerability:**
The `net/http/pprof` and `expvar` packages were imported using `_` in the main application entry points (`cmd/tesseract/gcp/main.go` and `cmd/tesseract/posix/main.go`). This automatically registers profiling (`/debug/pprof/*`) and metrics (`/debug/vars`) endpoints on the global `http.DefaultServeMux`. If an application uses `http.DefaultServeMux` or inadvertently exposes it, these endpoints can become publicly accessible, leading to a CWE-200 (Exposure of Sensitive Information to an Unauthorized Actor) vulnerability, which may reveal internal application state, memory profiles, and BadgerDB metrics.

**Learning:**
Anonymous imports (`_ "package"`) of debugging and profiling tools in production binaries introduce a significant risk by binding handlers to the default global multiplexer, which might be unintentionally exposed via public-facing servers.

**Prevention:**
Avoid importing `net/http/pprof` and `expvar` in production entry points. If profiling or metrics are required, explicitly register their handlers on a separate, internal-only `http.ServeMux` or use dedicated observability platforms (like OpenTelemetry) that do not rely on global HTTP handler registration.
