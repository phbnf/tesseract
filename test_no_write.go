package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
)

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Read the body, ignore error
        io.ReadAll(r.Body)
    })

    wrappedHandler := http.MaxBytesHandler(handler, 10)
    server := httptest.NewServer(wrappedHandler)
    defer server.Close()

    body := strings.Repeat("a", 20)
    resp, err := http.Post(server.URL, "text/plain", strings.NewReader(body))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Printf("Status code: %d\n", resp.StatusCode)
}
