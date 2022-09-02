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
        // these can be lower-cased b/c go converts them to their canonical form: https://github.com/golang/go/blob/2ebe77a2fda1ee9ff6fd9a3e08933ad1ebaea039/src/net/textproto/header.go#L34
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

