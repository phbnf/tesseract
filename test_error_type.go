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
        _, err := io.ReadAll(r.Body)
        if err != nil {
            fmt.Printf("Error type: %T\n", err)
            fmt.Printf("Error string: %s\n", err.Error())
        }
    })

    wrappedHandler := http.MaxBytesHandler(handler, 10)
    server := httptest.NewServer(wrappedHandler)
    defer server.Close()

    body := strings.Repeat("a", 20)
    http.Post(server.URL, "text/plain", strings.NewReader(body))
}
