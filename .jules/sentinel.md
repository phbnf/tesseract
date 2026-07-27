## 2024-06-29 - [Fix CWE-200 in Cloud Personalities]
**Vulnerability:** The `cmd/tesseract/gcp/main.go` and `cmd/tesseract/posix/main.go` files contained anonymous imports for `net/http/pprof` and `expvar`. This exposed the `/debug/pprof/*` and `/debug/vars` endpoints on the global `http.DefaultServeMux`, leading to CWE-200 Information Exposure.
**Learning:** These entry points should avoid using `net/http/pprof` and `expvar` because `http.DefaultServeMux` is generally exposed in HTTP server entry points.
**Prevention:** Cloud personalities utilize OpenTelemetry (otel) for handlers and deliberately avoid exposing `net/http/pprof` endpoints. Avoid anonymous imports for metrics and profiling on public-facing multiplexers.
