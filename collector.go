package main

import (
    "fmt"
    "log"
    "net/http"
    "golang.org/x/net/html"
)

func main() {
        log.Print("starting server...")
        http.HandleFunc("/", handler)

        port := "80"

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }
}

func getRequestIp(r *http.Request) string {
        headers := [3]string{"x-forwarded-for", "true-client-iP", "x-real-ip"}
        for _, header := range headers {
            ip := r.Header.Get(header)
            if ip != "" {
                return ip
            }
        }
        return "<unknown>"
}

func handler(w http.ResponseWriter, r *http.Request) {
        ip := getRequestIp(r)
        log.Printf("ip: %s", ip)
        log.Printf("path: %s", html.EscapeString(r.URL.Path))
        fmt.Fprint(w, "Hello World!\n")
}

