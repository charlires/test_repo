package main

import "fmt"

func greet() string {
    // Cambios introducidos en rama1
    return "Hola desde rama1"
}

func main() {
    fmt.Println(greet())
}