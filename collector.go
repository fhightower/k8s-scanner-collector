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

func handler(w http.ResponseWriter, r *http.Request) {
        log.Printf("path: %s", html.EscapeString(r.URL.Path))
        fmt.Fprint(w, "Hello World!\n")
}

