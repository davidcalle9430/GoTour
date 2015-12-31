package main

import(
  "fmt"
)


type MiFlotante float64

func (f MiFlotante) Abs() MiFlotante {
  if f < 0 {
    return MiFlotante(-f)
  }

  return f
}
func main() {
  f1 := MiFlotante(-5.0)
  f2 := MiFlotante(5)
  fmt.Println(f1.Abs() , " " , f2.Abs())
}
