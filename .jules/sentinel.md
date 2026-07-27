## 2026-06-06 - Prevent CWE-200 Information Exposure via DefaultServeMux
**Vulnerability:** Anonymous imports of `net/http/pprof` and `expvar` automatically register handlers on `http.DefaultServeMux`. If an application uses `http.DefaultServeMux` for public-facing servers without restriction, it inadvertently exposes sensitive profiling data and metrics (CWE-200).
**Learning:** These side-effect imports are often added for debugging and forgotten, leaking internal state to anyone who can reach the HTTP server.
**Prevention:** Avoid `_ "net/http/pprof"` and `_ "expvar"` in production binaries. If profiling or metric endpoints are needed, register them explicitly on a dedicated, internal-only `http.ServeMux` that is not exposed to the public network.
