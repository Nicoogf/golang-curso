package main

import (
	f "fmt"
)

type Persona struct{
	Nombre string
}

func Cambiame ( p Persona ){
	p.Nombre = "Nombre editado"
}

func main(){
 ejemplo := Persona{Nombre: "Nicolas"}

 Cambiame(ejemplo)
 f.Println(  (ejemplo).Nombre )
}