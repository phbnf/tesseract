## 2024-05-24 - Unintentional pprof Exposure
**Vulnerability:** Profiling endpoints were automatically exposed on `/debug/pprof` by importing `_ "net/http/pprof"`.
**Learning:** Importing `net/http/pprof` registers its handlers on the global `http.DefaultServeMux`. If an application exposes this mux without authentication, it inadvertently exposes sensitive profiling data.
**Prevention:** Avoid anonymous imports of `net/http/pprof` in production-facing applications or explicitly restrict access to profiling endpoints.
