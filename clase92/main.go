package main

import (
	f "fmt"
)


func main() {
	arreglo := []int { 5,5,10,25,50,5 }
	sumaDeArreglo := sumaDeValores(arreglo...)

	f.Println(sumaDeArreglo)
}

func sumaDeValores ( x ...int ) int {
	resultado := 0
	for _ , v := range x {
		resultado += v
	}
	return resultado
}