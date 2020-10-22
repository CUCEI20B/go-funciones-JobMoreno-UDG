package main

import "fmt"
func Burbuja(ListaDesordenada []int64) {
 var auxiliar int64
 for i := 0; i < len(ListaDesordenada); i++ {
  for j := 0; j < len(ListaDesordenada); j++ {
   if ListaDesordenada[i] < ListaDesordenada[j] {
    auxiliar = ListaDesordenada[i]
    ListaDesordenada[i] = ListaDesordenada[j]
    ListaDesordenada[j] = auxiliar
   }
  }
 }
}
func main(){
	a := []int64{1, 4, 2}
  Burbuja(a)
}