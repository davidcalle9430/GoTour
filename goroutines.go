package main


import (
  "fmt"
  "time"
)

func say(what string){
  for i := 0 ; i < 5 ; i++ {
    time.Sleep(10 * time.Millisecond)
    fmt.Println(what)
  }
}

func main() {
  go say("Hello")
  go say("bye")
  time.Sleep(1000 * time.Millisecond)
}
