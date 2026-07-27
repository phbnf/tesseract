## 2025-07-05 - [HIGH] Fix CWE-200 profiling and metrics endpoints exposure
**Vulnerability:** Profiling (`net/http/pprof`) and metrics (`expvar`) endpoints were exposed via anonymous imports on `http.DefaultServeMux` in the `cmd/tesseract/gcp/main.go` and `cmd/tesseract/posix/main.go` binaries. This could lead to CWE-200 Information Exposure.
**Learning:** Anonymous imports of `net/http/pprof` or `expvar` automatically register handlers on `http.DefaultServeMux`. If `http.ListenAndServe()` is used, this exposes these endpoints on any interface unless specifically restricted or handled correctly on a custom multiplexer.
**Prevention:** Avoid anonymous imports of `net/http/pprof` or `expvar` in production binaries to prevent inadvertent exposure of diagnostic endpoints.
