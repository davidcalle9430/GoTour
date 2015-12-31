package main

import(
  "fmt"
  "io"
  "strings"
)

func main() {
  r := strings.NewReader("Hola! WOOO")

  b := make([]byte, 8)

  for{
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v",  n , err, b)
    fmt.Println("")
    fmt.Printf("b[:n] = %q",b[:n])
    if err == io.EOF {
      break
    }
  }
}
