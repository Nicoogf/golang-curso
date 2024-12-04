package main

import (
	f "fmt"
)

func main() {
	var primer_arreglo [5]int
	primer_arreglo[0] = 1
	primer_arreglo[1] = 2
	primer_arreglo[2] = 3
	primer_arreglo[3] = 4
	primer_arreglo[4] = 5
  
	for indice , valor := range primer_arreglo{
		f.Println(indice ," ",valor)
	}

}