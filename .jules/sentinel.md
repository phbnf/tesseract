
## 2024-05-11 - Prevent Information Exposure via pprof and expvar Endpoints (CWE-200)
**Vulnerability:** Anonymous imports of `net/http/pprof` and `expvar` automatically exposed profiling (`/debug/pprof/*`) and metrics (`/debug/vars`) endpoints on the global `http.DefaultServeMux`. These endpoints can reveal sensitive application internals and potential memory/CPU consumption patterns to attackers.
**Learning:** Default behavior of these packages directly modifies the global state, making it surprisingly easy to unintentionally expose internal diagnostics in production binaries.
**Prevention:** Avoid anonymous imports of `net/http/pprof` and `expvar` in production entry points. If profiling is needed, register endpoints explicitly on an internal, authenticated multiplexer instead of the globally exposed `http.DefaultServeMux`.
