## 2024-05-27 - Remove exposed pprof and expvar endpoints
**Vulnerability:** Automatically exposed profiling endpoints (`/debug/pprof/*`) and variables (`/debug/vars`) via anonymous imports of `net/http/pprof` and `expvar` (CWE-200).
**Learning:** Anonymous imports of these packages automatically bind to `http.DefaultServeMux`, which may unintentionally expose sensitive performance and environment data on the primary HTTP listener.
**Prevention:** Avoid anonymous imports of `net/http/pprof` and `expvar`. If profiling is needed, explicitly configure and serve it on a separate, protected listener/port.
