package main

import (
	"encoding/json"
	f "fmt"
	
)

type Persona struct{
	Nombre string
	Apellido string
	Edad int
}

func main(){
	p1 := Persona{Nombre: "Nicolas" ,Apellido: "Gonzales" ,Edad: 28}
	p2 := Persona{Nombre: "Adrian" ,Apellido: "Nario" ,Edad: 48}

	personas := []Persona {p1,p2}

	f.Println(personas)

	bs , err := json.Marshal(personas)

	if err != nil {
		f.Println(err)
	}

	f.Println(string(bs))
} 