## 2026-03-17 - [Sentinel] Fix CWE-200
**Vulnerability:** Profiling endpoints automatically exposed on HTTP default serve mux (CWE-200) due to importing _ "net/http/pprof" in public-facing custom multiplexers in AWS/GCP binaries.
**Learning:** When mitigating global http.DefaultServeMux exposure, be careful not to inadvertently import and expose net/http/pprof on public-facing custom multiplexers, as this introduces CWE-200 (profiling endpoint exposure).
**Prevention:** Avoid importing _ "net/http/pprof" in public-facing HTTP server binaries to prevent inadvertent exposure of profiling endpoints.
