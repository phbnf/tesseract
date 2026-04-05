## 2026-04-05 - Remove exposed profiling endpoints
**Vulnerability:** Profiling (`net/http/pprof`) and metrics (`expvar`) endpoints were exposed via anonymous imports. This can lead to a CWE-200 vulnerability, leaking sensitive application data.
**Learning:** Anonymous imports of `net/http/pprof` or `expvar` automatically register HTTP handlers on the default serve mux `http.DefaultServeMux`. If `http.ListenAndServe` or `http.Server.ListenAndServe` uses the default serve mux without overriding handlers, the debug endpoints become publicly available.
**Prevention:** Avoid blank imports of `net/http/pprof` and `expvar` in binaries meant for production that expose an HTTP server. Only expose diagnostic endpoints internally or over authenticated interfaces.
