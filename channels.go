package main

import "fmt"

func sum(a []int, c chan int){
  sum := 0
  for _ , v := range a {
    sum += v
  }
  fmt.Println("El resultado es", sum)
  c <- sum // enviar sum a c
}

func main() {
  a := []int{1,1,1,2,2,2}
  c := make(chan int)

  go sum( a[:len(a)/2] , c )
  go sum( a[len(a)/2:] , c )
  go sum( a[:len(a)/2] , c )
  go sum( a[len(a)/2:] , c )
  go sum( a[:len(a)/2] , c )
  go sum( a[len(a)/2:] , c )
  go sum( a[:len(a)/2] , c )
  go sum( a[len(a)/2:] , c )
  go sum( a[:len(a)/2] , c )
  go sum( a[len(a)/2:] , c )
  x, y, z, w := <-c , <-c, <-c, <-c // recibo de c
  fmt.Println(x, y, z, w)
}
