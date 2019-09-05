package main

import "fmt"

func menor(a, b int) int{
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Printf("O menor é %d\n", menor(1,2))
    fmt.Printf("O menor é %d\n", menor(2,1))
}
