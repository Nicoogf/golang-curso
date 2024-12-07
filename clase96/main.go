package main

import (
	f "fmt"
)

func main() {
	x := practica()
	
	valor := x()

	f.Println(valor)
}


func practica() func() int {
	return func() int {
		return 1942
	}
}