package main

import (
    "errors"
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
            var maxBytesErr *http.MaxBytesError
            if errors.As(err, &maxBytesErr) {
                fmt.Println("Caught http.MaxBytesError")
                w.WriteHeader(http.StatusRequestEntityTooLarge)
                return
            }
        }
        w.WriteHeader(http.StatusOK)
    })

    // Manually wrapping with MaxBytesReader, similar to what http.MaxBytesHandler does internally for Body
    // But http.MaxBytesHandler ALSO wraps the ResponseWriter to return 413 if Write is called.
    // However, we are reading the body.

    // Let's test http.MaxBytesHandler specifically.
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
