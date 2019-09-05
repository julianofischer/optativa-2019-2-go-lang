package main

import "fmt"

/*
 Escreva uma função que receba dois números inteiros e retorne o menor 
 entre eles.
*/
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
