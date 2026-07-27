## 2025-02-15 - Exposure of profiling endpoints (CWE-200)
**Vulnerability:** The posix and gcp cloud personalities were exposing `net/http/pprof` and `expvar` on the default HTTP multiplexer via anonymous imports.
**Learning:** Anonymous imports of `net/http/pprof` and `expvar` automatically register profiling and metrics endpoints (`/debug/pprof/*` and `/debug/vars`) on the globally accessible `http.DefaultServeMux`, leading to CWE-200.
**Prevention:** Do not use anonymous imports for `net/http/pprof` or `expvar` in production applications. If profiling is needed, register handlers manually on an internal-only multiplexer (on a separate port).
