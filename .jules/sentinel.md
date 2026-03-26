## 2025-05-28 - [CWE-200] Prevent Profiling Endpoint Exposure in HTTP Servers

**Vulnerability:** Profiling endpoints were inadvertently exposed globally on `http.DefaultServeMux` by importing `_ "net/http/pprof"` in `cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go`. This exposed potentially sensitive memory and CPU profiling data (CWE-200) on endpoints that might be publicly accessible.

**Learning:** The `net/http/pprof` package automatically registers handlers on `http.DefaultServeMux` upon initialization (`init()`). Any HTTP server that uses this mux or relies on default HTTP handling behavior without isolating its routing can inadvertently expose these debugging endpoints.

**Prevention:** Never import `_ "net/http/pprof"` in production entry points unless access to `http.DefaultServeMux` is strictly restricted (e.g., bound only to `localhost` on an internal admin port). If profiling is necessary, explicitly register the handlers on a separate, securely bound `http.ServeMux`. Use OpenTelemetry or similar dedicated observability pipelines, as preferred by this project.
