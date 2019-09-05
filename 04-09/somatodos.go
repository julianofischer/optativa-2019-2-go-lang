package main

import "fmt"

func soma(nums ...int) int {
    total := 0
    for _, num := range nums {
        total = total + num
    }
    return total
}

func main() {
    a := []int{1,2,3,4}
    fmt.Printf("A soma é %d\n", soma(1,2))
    fmt.Printf("A soma é %d\n", soma(2,1,3))
    fmt.Printf("A soma é %d\n", soma(a...))
}
