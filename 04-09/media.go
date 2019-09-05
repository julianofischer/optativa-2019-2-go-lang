package main

import "fmt"

func soma(nums ...int) float64 {
    total := 0.0
    length := float64(len(nums))
    for _, num := range nums {
        total = total + float64(num)
    }
    return total/length
}

func main() {
    a := []int{1,2,3,4}
    fmt.Printf("A média é %f\n", soma(1,2))
    fmt.Printf("A média é %f\n", soma(2,1,3))
    fmt.Printf("A média é %f\n", soma(a...))
}
