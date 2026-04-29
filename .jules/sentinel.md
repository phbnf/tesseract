## 2025-10-24 - Exposed Profiling Endpoint (CWE-200) in posix personality
**Vulnerability:** The `posix` personality exposed Go profiling (`net/http/pprof`) and `expvar` variables globally on the default `http.ServeMux` due to anonymous imports (`_ "net/http/pprof"` and `_ "expvar"`).
**Learning:** Anonymous imports of diagnostic packages bind to the default global multiplexer, exposing sensitive internal state, memory details, and potentially allowing DoS attacks on production endpoints unless explicitly secured or firewalled.
**Prevention:** Avoid anonymous imports of `net/http/pprof` or `expvar` in production binaries. If profiling is needed, explicitly mount it on an internal, authenticated, or separate administrative port, rather than the primary public-facing HTTP server.
