package main 

import (
	f "fmt"
)

type persona struct{
	Nombre string
	Apellido string 
	Edad int
}

func (p persona) presentar(){
	f.Printf("Buenos dias , me llamo %s %s" , p.Nombre , p.Apellido)
}

func main() {
	p1 := persona{Nombre: "Nicolas" , Apellido: "Gonzales" , Edad: 28,}
	p1.presentar()

}