## 2026-04-09 - Fix CWE-200 profiling endpoint exposure on posix backend
**Vulnerability:** The POSIX cloud personality entry point (`cmd/tesseract/posix/main.go`) included blank imports for `net/http/pprof` and `expvar`.
**Learning:** Including these blank imports automatically registers profiling (`/debug/pprof/*`) and metrics (`/debug/vars`) endpoints on the global `http.DefaultServeMux`. If `http.DefaultServeMux` is used or exposed in a public-facing application, this leaks sensitive internal state and performance data (CWE-200).
**Prevention:** Avoid blank imports of `net/http/pprof` and `expvar` in production binaries. Use OpenTelemetry (otel) for observability, as done in the AWS and GCP handlers, and deliberately avoid exposing these default endpoints.
