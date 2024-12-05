package main

import (
	f "fmt"
)

type persona struct{
	nombre string
	apellido string
	edad int
}

type agenteSecreto struct{
	persona
	lpm bool
}
	

func main() {
	as := agenteSecreto{
		persona: persona{nombre: "Juan" , apellido: "Gonzales",edad: 45},
		lpm: true,
	}

	f.Println(as)
}