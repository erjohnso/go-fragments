package main

import (
    "net/http"
    "log"
    "fmt"
)

func main() {
    port := 8080

    http.Handle("/", http.FileServer(http.Dir(".")))
    log.Printf("=> Listening on port %d\n", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
    if err != nil {
        log.Fatal("ListenAndServe(\":8080\"):", err)
    }
}
