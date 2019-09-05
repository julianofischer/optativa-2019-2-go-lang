package main

import "fmt"

/*
 * Escreva uma função que receba dois números inteiros e retorne dois 
 * números: menor, maior
*/
func menor(a, b int) (int, int){
    if a < b {
        return a, b
    }
    return b, a
}

func main() {
    men, maior := menor(1,2)
    fmt.Printf("O menor é %d e o maior é %d\n", men, maior)
    men, maior = menor(2,1)
    fmt.Printf("O menor é %d e o maior é %d\n", men, maior)
}
