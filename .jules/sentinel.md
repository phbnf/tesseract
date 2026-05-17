## 2026-03-16 - Prevent Profiling Endpoint Exposure
**Vulnerability:** Automatic exposure of `net/http/pprof` endpoints (`/debug/pprof`) via anonymous import.
**Learning:** In Go, anonymously importing `net/http/pprof` automatically registers profiling handlers on `http.DefaultServeMux`. If a public HTTP server uses the default mux, it exposes sensitive application internals (memory, CPU, command-line arguments, goroutines) and creates a potential Denial of Service (DoS) and Information Leakage (CWE-200) vector.
**Prevention:** Avoid anonymous imports of `net/http/pprof` in production-facing binaries (like AWS/GCP/POSIX cloud personalities). If profiling is needed, register it explicitly on a separate, internal-only `ServeMux` protected by authentication or a private network interface.
