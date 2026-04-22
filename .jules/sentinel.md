## 2024-05-20 - Remove auto-exposed pprof and expvar endpoints
**Vulnerability:** pprof profiling endpoints and expvar metrics were automatically exposed on `/debug/pprof` and `/debug/vars` globally because the `cmd/tesseract/posix/main.go` entrypoint anonymously imported `net/http/pprof` and `expvar` (CWE-200).
**Learning:** Anonymous imports of these packages automatically bind their handlers to `http.DefaultServeMux`. If the application uses the default mux, these endpoints are exposed to the internet.
**Prevention:** Avoid anonymous imports of `net/http/pprof` and `expvar`. If profiling or metrics are needed, explicitly bind them to a separate internal server or use a framework like OpenTelemetry.
