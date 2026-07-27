## 2026-06-14 - Prevent CWE-200 Information Exposure via Anonymous HTTP Handlers

**Vulnerability:**
The `cmd/tesseract/posix/main.go` and `cmd/tesseract/gcp/main.go` files included anonymous imports for `net/http/pprof` and `expvar`. These imports automatically bind their handlers to `http.DefaultServeMux`. If `http.DefaultServeMux` is exposed, this causes CWE-200 (Information Exposure), revealing sensitive profiling and metric information publicly without authentication.

**Learning:**
Anonymous imports of diagnostic or metric libraries such as `_ "net/http/pprof"` or `_ "expvar"` are dangerous because they mutate global state (`http.DefaultServeMux`) implicitly. In production or cloud personalities, relying on OpenTelemetry or other explicit metric bindings is preferred, avoiding unintended endpoint exposure.

**Prevention:**
Avoid `_ "net/http/pprof"` or `_ "expvar"` imports in binary entry points (`main.go`). If profiling or metric endpoints are needed, bind them explicitly to an internal, secured mux that is not exposed to public traffic. Check for similar issues using static analysis scanners and code reviews targeting anonymous imports.
