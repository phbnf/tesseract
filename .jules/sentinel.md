## 2024-05-30 - Prevent pprof endpoint exposure on public servers
**Vulnerability:** Importing _ "net/http/pprof" or _ "expvar" automatically registers debug endpoints on http.DefaultServeMux, which can lead to information disclosure if the default mux is exposed on public HTTP servers.
**Learning:** Default multiplexers are risky. Cloud personalities entry points (cmd/tesseract/*/main.go) expose http.Server without explicitly configuring a dedicated mux, causing http.DefaultServeMux to be used with the registered debug endpoints exposed publicly.
**Prevention:** Avoid blank imports of net/http/pprof and expvar in production application entry points unless bound to an internal-only port or a specific mux.
