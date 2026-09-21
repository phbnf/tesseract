## 2026-03-27 - Exposed Profiling Endpoint
**Vulnerability:** The standard library `net/http/pprof` automatically registers profiling endpoints to `http.DefaultServeMux`, exposing sensitive system information.
**Learning:** In the POSIX server entrypoint, this package was anonymously imported (`_ "net/http/pprof"`), leading to CWE-200.
**Prevention:** Ensure that `net/http/pprof` is not imported, especially on public-facing custom multiplexers.
