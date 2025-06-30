package main

import (
    "fmt"
    "net/http"
)

func holaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hola, mundo!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", holaHandler)
    fmt.Printf("Server listening on :%s\n", port)
    http.ListenAndServe(fmt.Sprintf(":%s"port), nil)
}
