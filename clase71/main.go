package main

import(
	f "fmt"
)

func main() {
	arreglo_dos := [12]int {40,41,42,43,44,45,46,47,48,49,50,51}

	f.Println(arreglo_dos[2:7])
	f.Println(arreglo_dos[7:])
	f.Println(arreglo_dos[4:9])
	f.Println(arreglo_dos[3:8])
}
