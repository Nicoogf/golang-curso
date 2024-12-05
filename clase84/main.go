package main

import (
	f "fmt"
)
	

func main() {
	as := struct{
			nombre string
			apellido string
			edad int
		}{nombre: "Juan" , apellido: "Gonzales",edad: 45}
	

	f.Println(as)
}