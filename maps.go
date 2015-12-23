package main

import(
  "fmt"
)

type Vertex struct{
  Lat, Lon float64
}


var m map[string] Vertex

func main() {
  m = make(map[string]Vertex)
  m["David"] = Vertex{ 10.6, 11.21343 }
  fmt.Println(m["David  "])
}
