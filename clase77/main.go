package main

import (
	f "fmt"
)

func main() {

	slice_uno := []string {"Batman", "Jefe" , "Vestido de Negro"}
	slice_dos := []string {"Robin", "Ayudante" , "Vestido de Colores"}
	x := [][]string { slice_uno , slice_dos }
	f.Println(x)

	for valor, indice := range x {
		f.Println("Indice " , valor , "Array : " , indice)
	}

}