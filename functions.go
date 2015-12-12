package main

import(
  "fmt"
)

const PI = 3.14
func add(x int, y int)  int  {
  return x + y
}
/*
* Se omite el el tipo de los parámetros excepto el del último
*/
func simplifiedAdd( x , y , z int) int{
  return x + y + z
}

func swap( x , y string ) (string, string) {
  return y , x
}

func split(sum int) ( x , y int ){
  x = sum * 4 / 9
  y = sum - x
  return
}
func main(){
  fmt.Println("El resultado de la suma es: ")
  fmt.Println(add(2,5))
  fmt.Println(simplifiedAdd(1,2,3))
  hello , world := swap("hello", "world")
  fmt.Println( hello + "-" + world )
  first , second := split(5)
  fmt.Println(first, second)
  var c , python , java bool
  var i int
  fmt.Println(i , c , python , java)
  var a , b , d int = 1 , 2 , 3
  fmt.Println(a , b , d)
  e := 14
  fmt.Println(e)
  fmt.Println(float64(e))
  f := -1
  fmt.Println(uint(f))
  fmt.Println(PI)
}
