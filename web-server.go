package main

import(
  "fmt"
  "net/http"
)

type Hello struct{}
type Persona struct{
  Saludo string
  Nombre string
}


func (p Persona) ServeHTTP (w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, p.Saludo)
}
func (h Hello) ServeHTTP (w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, " Hola! ")
}

func main() {
  persona := Persona{ "David", "Saludos"}
  http.Handle("/saludo/", persona)
  fmt.Println("llega")
  go http.ListenAndServe("localhost:4000", nil)
  fmt.Println("llega")
}
