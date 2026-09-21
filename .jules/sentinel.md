## 2024-06-07 - CWE-200: Profiling Endpoint Exposure

**Vulnerability:** The `net/http/pprof` and `expvar` packages were imported in the production binaries (`cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go`). These packages automatically register endpoints like `/debug/pprof/*` and `/debug/vars` on the global `http.DefaultServeMux`, which is exposed to the internet.

**Learning:** This exposes sensitive internal application state and profiling endpoints to anyone who can access the public server, which is a significant Information Exposure vulnerability (CWE-200).

**Prevention:** Never import `net/http/pprof` or `expvar` in production entry points without robust authentication/authorization. Rely on internal monitoring solutions (like OpenTelemetry, which is already used in this codebase) instead.
