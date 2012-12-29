package main

import (
    "net/http"
    "log"
    "fmt"
)

type LoggingRequest struct {
    path string
}

func LogAndServe(w http.ResponseWriter, req *http.Request) {
    log.Printf("[%s] \"%s %s %s\"\n", req.RemoteAddr, req.Method, req.URL, req.Proto)
    http.FileServer(http.Dir("."))
}

//func myFileServer(c *http.Conn, req *http.Request) {
//    log.Printf("[%s] \"%s %s %s\"\n", req.RemoteAddr, req.Method, req.URL, req.Proto)
//    http.FileServer(http.Dir("."))
//}

func main() {
    port := 8080

    loggedRequest := new(LoggingRequest)
    loggedRequest.path = "."
    //http.Handle("/", http.FileServer(http.Dir(".")))
    http.Handle("/", http.HandlerFunc(LogAndServe))
    log.Printf("=> Listening on port %d\n", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
    if err != nil {
        log.Fatal("ListenAndServe(\":8080\"):", err)
    }
}
