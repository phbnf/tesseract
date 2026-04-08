## 2025-04-08 - Remove pprof and expvar from public endpoints
**Vulnerability:** Publicly accessible profiling (`/debug/pprof/*`) and metrics (`/debug/vars`) endpoints due to importing `_ "net/http/pprof"` and `_ "expvar"`.
**Learning:** Importing these packages registers endpoints on `http.DefaultServeMux`. If a public server falls back to this default mux (e.g., via `http.Handle("/", handler)` without specifying a custom mux), these endpoints are exposed, creating a CWE-200 vulnerability.
**Prevention:** Avoid using `_ "net/http/pprof"` and `_ "expvar"` in production binaries. Explicitly map these to a private/internal server if needed, or use safe observability libraries like OpenTelemetry instead.
