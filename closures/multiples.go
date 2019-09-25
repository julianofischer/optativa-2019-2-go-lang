package main

import "fmt"

func generator(n int) func() int {
    item := 0
    return func() int {
        item = item + 1
        return n*item
    }
}

func main() {
   next := generator(2)
   fmt.Println(next())
   fmt.Println(next())
   fmt.Println(next())
   fmt.Println(next())
}
