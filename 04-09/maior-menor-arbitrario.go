package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func menormaior(nums ...int) (int, int) {
    var maior int = MinInt
    var menor int = MaxInt

    for _, num := range nums {
        if num > maior {
            maior = num
        }

        if num < menor{
           menor = num
        }
    }
    return menor, maior
}

func main() {
    a := []int{-5,1,-10,2,3,4,205}
    menor, maior := menormaior(a...)
    fmt.Printf("%d, %d\n", menor, maior)
}
