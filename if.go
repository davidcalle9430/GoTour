package main

import(
  "fmt"
  "math"
  "runtime"
)

func sqrt(x float64) string  {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux. ")
  default:
    fmt.Println("%s.", os)
  }
}
