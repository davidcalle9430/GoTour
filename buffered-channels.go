package main

import(
  "fmt"
  "time"
)

func main() {
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  //ch <- 3
  //ch <- 4
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  time.Sleep(1000 * time.Millisecond)
}
