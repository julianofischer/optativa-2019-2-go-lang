package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
