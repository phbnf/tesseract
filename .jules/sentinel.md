## 2024-05-24 - CWE-200 Profiling Endpoint Exposure
**Vulnerability:** Import of `_ "net/http/pprof"` and `_ "expvar"` in binary entry points registers debugging endpoints (`/debug/pprof/*` and `/debug/vars`) on the `http.DefaultServeMux`.
**Learning:** This exposes sensitive application internals (metrics, memory profiles, CPU profiles) that an attacker could use to glean information or conduct denial of service attacks.
**Prevention:** Avoid registering global side-effect imports like `_ "net/http/pprof"` and `_ "expvar"` in binaries that start HTTP servers. Instead, attach these endpoints explicitly to a dedicated, restricted-access multiplexer if profiling is necessary.
