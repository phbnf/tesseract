## 2024-05-18 - [Fix profiling endpoint exposure (CWE-200)]
**Vulnerability:** Unintentional exposure of the `net/http/pprof` profiling endpoint on `http.DefaultServeMux`.
**Learning:** The `_ "net/http/pprof"` import implicitly attaches its debug handlers (e.g., `/debug/pprof/`) to the default HTTP multiplexer (`http.DefaultServeMux`). If `http.DefaultServeMux` is directly or indirectly exposed, an attacker can access sensitive runtime information.
**Prevention:** Do not use `_ "net/http/pprof"` in production endpoints unless heavily restricted (e.g., authentication or an isolated internal port). The project uses OpenTelemetry for instrumentation instead.
