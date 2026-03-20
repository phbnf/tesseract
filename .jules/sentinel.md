## 2025-03-20 - [Remove net/http/pprof to prevent profiling endpoint exposure]
**Vulnerability:** Profiling endpoint is automatically exposed on /debug/pprof because of `_ "net/http/pprof"` import.
**Learning:** Importing `_ "net/http/pprof"` automatically registers profiling handlers on `http.DefaultServeMux`, which exposes them on the main HTTP server without authentication.
**Prevention:** Avoid importing `_ "net/http/pprof"` in production binaries. If profiling is needed, register pprof handlers on a separate, internal port or add authentication.
