package main

import (
	f "fmt"
)

func main() {
 x := foo(1,23,47,5)
 f.Println("El valor de x fue " , x )
}


func foo ( x ...int) int {
	f.Printf("El tipo es %T" , x)

	suma := 0
	for i , v := range x {
		suma += v
		f.Println("El indice fue " , i , "su valor fue " , v)
	}

	f.Println(suma)
	return suma

}