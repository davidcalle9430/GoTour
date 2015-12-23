package main

import(
  "fmt"
  "math"
)
type Vertex struct{
  Lat, Lon float64
}

func (v *Vertex) Abs() float64{
  return math.Sqrt((v.Lat * v.Lat) + (v.Lon * v.Lon))
}
func main() {
  vertex := Vertex { 3 , 4}
  fmt.Println(vertex.Abs())
}
