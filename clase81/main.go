package main

import (
	f "fmt"
)

type Persona struct{
	nombre string
	apellido string
	edad int
}

func main() {
	var nicolas Persona 

	nicolas.nombre = "Juan"
	nicolas.apellido = "Gomez"
	nicolas.edad = 25

	f.Println(nicolas)

	usuario_dos := Persona {
		nombre : "Carlos",
		apellido: "Gonzales",
		edad: 48,
	}

	f.Println(usuario_dos)
}