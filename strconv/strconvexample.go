package main

import "fmt"
import "strconv"

func main() {
    var d int64
    var f float64
    d, _ = strconv.ParseInt("18", 10, 64)
    f, _ = strconv.ParseFloat("18.8", 64)
    fmt.Printf("Converted int: %d\n", d)
    fmt.Printf("Converted float: %f\n", f)
    d32,_ := strconv.Atoi("13")
    fmt.Printf("Another converted int %d\n", d32)
    s := strconv.Itoa(50)
    fmt.Printf("What abou converting integer to string? %s\n", s)

}

