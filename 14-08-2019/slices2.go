package main

import "fmt"
import "reflect"

func main() {

    var s []int = []int{1, 2, 3, 4, 5}
    fmt.Printf("The type of s is %s\n", reflect.TypeOf(s))
    var a [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Printf("The type of a is %s\n", reflect.TypeOf(a))
    printSlice(s)

    //compiling error -> expecting slice
    //printSlice(a)
}

func printSlice(s []int){
   fmt.Printf("%#v", s)
}
