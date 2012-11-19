package main

import (
    "net/http"
    "log"
    "fmt"
)

//func myFileServer(c *http.Conn, req *http.Request) {
//    log.Printf("[%s] \"%s %s %s\"\n", req.RemoteAddr, req.Method, req.URL, req.Proto)
//    http.FileServer(http.Dir("."))
//}
//
//func main() {
//    port := 8080
//
//    http.Handle("/", http.HandlerFunc(myFileServer))
//    log.Printf("=> Listening on port %d\n", port)
//
//    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
//    if err != nil {
//        log.Fatal("ListenAndServe(\":8080\"):", err)
//    }
//}


func main() {
    port := 8080

    http.Handle("/", http.FileServer(http.Dir(".")))
    log.Printf("=> Listening on port %d\n", port)
    err := http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
    if err != nil {
        log.Fatal("ListenAndServe(\":8080\"):", err)
    }
}
