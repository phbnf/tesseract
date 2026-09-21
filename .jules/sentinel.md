## 2026-05-30 - Exposure of Profiling and Metrics Endpoints (CWE-200)
**Vulnerability:** The GCP and POSIX entry points in `cmd/tesseract` were importing `_ "net/http/pprof"` and `_ "expvar"`. This caused profiling and metric endpoints (`/debug/pprof/*` and `/debug/vars`) to be automatically registered and exposed on `http.DefaultServeMux`.
**Learning:** Anonymous imports of standard library debugging tools can silently expose sensitive internal endpoints to the internet if `http.DefaultServeMux` is bound to a public interface.
**Prevention:** Avoid using anonymous imports for debugging packages in production binaries, especially when creating custom server multiplexers, and always verify what routes are registered on the public listener.
