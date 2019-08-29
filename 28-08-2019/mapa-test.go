package main
import "fmt"

func main(){
    m := make(map[int]string)
    m[0] = "Juliano"
    value, exists := m[0]
    fmt.Printf("%s %v\n",value,exists)

    fmt.Println(m[0])
}
