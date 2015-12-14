package main

import(
  "fmt"
)

func main() {
  defer fmt.Println(", Mundo!")
  fmt.Println("Hola")


  fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
