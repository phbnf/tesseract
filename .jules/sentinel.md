## 2024-05-18 - CWE-200 Information Exposure via Profiling and Metrics Endpoints
**Vulnerability:** The `cmd/tesseract/posix/main.go` entry point imported `_ "net/http/pprof"` and `_ "expvar"`, which inadvertently exposed profiling and metrics endpoints (`/debug/pprof/*` and `/debug/vars`) on the `http.DefaultServeMux`, leading to potential CWE-200 vulnerabilities.
**Learning:** Cloud personality entry points (e.g., posix, aws, gcp) must avoid importing these packages to prevent exposing sensitive internal state and performance metrics on public-facing HTTP servers.
**Prevention:** Do not import `net/http/pprof` or `expvar` in production entry points. Use secure, authenticated, and dedicated metric solutions (like OpenTelemetry) instead of attaching these to the default multiplexer.
