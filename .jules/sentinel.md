## 2025-02-23 - Prevent Profile and Debug Endpoint Exposure (CWE-200)

**Vulnerability:** Anonymous imports of `net/http/pprof` and `expvar` in HTTP server entrypoints (e.g., `cmd/tesseract/gcp/main.go`, `cmd/tesseract/posix/main.go`) automatically register endpoints on the global `http.DefaultServeMux`. If `http.DefaultServeMux` is unintentionally exposed or used by public-facing multiplexers without proper routing filters, it leaks sensitive debugging metrics and profiling data.

**Learning:** Blank imports like `_ "net/http/pprof"` mutate global state during initialization, which often goes unnoticed when bringing up HTTP servers, thereby exposing unintended attack surface on default HTTP handler setups.

**Prevention:** Avoid using anonymous imports for `pprof` or `expvar` in production entrypoints. Instead, if metrics are required, selectively register them onto internal, restricted, or administrative muxes explicitly rather than relying on global side-effects.
