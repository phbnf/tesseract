## 2024-05-24 - Profiling endpoint exposure (CWE-200)
**Vulnerability:** The public-facing posix server imported `_ "net/http/pprof"` and `_ "expvar"`.
**Learning:** Importing these packages automatically registers handlers on the default `http.DefaultServeMux`. If `http.DefaultServeMux` is used or exposed (e.g. implicitly, or if a custom server doesn't explicitly restrict routes), this exposes sensitive profiling data and internal metrics on endpoints like `/debug/pprof/*` and `/debug/vars`, leading to information disclosure.
**Prevention:** Avoid anonymous imports of `net/http/pprof` or `expvar` in binaries intended for production deployment without explicitly configuring a dedicated, secured (e.g. internal only) serve mux for debugging and profiling.
