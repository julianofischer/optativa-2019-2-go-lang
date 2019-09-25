package main

import (
    "fmt"
    "math/rand"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func randNumber() func() int {
    item := 0
    list := makeRange(1,60)
    return func() int {
        if item == 6 {
            item = 1
            list = makeRange(1,60)
        }
        item++
        r := rand.Intn(len(list))
        ret := list[r]
        list[r] = list[len(list)-1]
        list = list[:len(list)-1]
        return ret
    }
}

func main() {
    //fmt.Println(makeRange(1,61))
    gen := randNumber()
    i:=1
    for i<50 {
        fmt.Println(gen())
        i++
    }
}
