package main

import(
  "fmt"
  "math"
)

type Abser interface{
  Abs() float64
}

type MyFloat float64 // declaro mi tipo de flotante

func (f MyFloat) Abs() float64 { // le agrego el método de la intefaz
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
  var a Abser
  // declaro una variable de tipo interfazx
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
  // a MyFloat implements Abser
	a = &v
  // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// anxd does NOT implement Abser.
	//a = v 

	fmt.Println(a.Abs())
}
