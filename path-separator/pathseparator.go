package main

import "path"
import "fmt"

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println(dir, file)
}
