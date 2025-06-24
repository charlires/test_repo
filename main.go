package main

import "fmt"

func greet() string {
    // Cambios introducidos en rama2
    return "Hola desde rama2"
}

func main() {
    fmt.Println(greet())
}