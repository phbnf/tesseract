## 2024-03-02 - Unauthenticated Profiling Endpoints
**Vulnerability:** Anonymous import of `_ "net/http/pprof"` automatically registered profiling handlers on `http.DefaultServeMux` in the `tesseract/posix` server, which uses `http.Handle("/", logHandler)`.
**Learning:** This exposed potentially sensitive information (memory, CPU, and goroutine profiles) on the main public server port without authentication.
**Prevention:** Avoid `_ "net/http/pprof"` in production entry points unless `http.DefaultServeMux` is explicitly isolated from public traffic or protected by authentication middleware.
