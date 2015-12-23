package main

import(
  "fmt"
)

func main() {
  a := make([]int, 5)
  printSlice("a", a)
  b := make([]int, 0 , 5)
  printSlice("b", b)
  a = append(a, 4, 5)
  printSlice("a", a)

  for index, value := range a {
    fmt.Println("Valor ", value, " Indice ", index)
  }
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s, len(x), cap(x), x)

}
