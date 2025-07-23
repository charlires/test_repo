package main

import (
    "fmt"
    "net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "bar")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }
    http.HandleFunc("/foo", fooHandler)
    http.HandleFunc("/bar", barHandler)
    fmt.Printf("Server listening on :%s\n", port)
    http.ListenAndServe(fmt.Sprintf(":%s"port), nil)
}
