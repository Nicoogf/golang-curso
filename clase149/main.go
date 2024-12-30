package main

import (
	"fmt"
)

type persona struct{
	nombre string
}

func(p *persona) hablar() {
	fmt.Println("Hola soy  : " , p.nombre)
}

type humano interface {
	hablar()
}

func DiAlgo(h humano) {
	h.hablar()
}


func main(){
	p1 := persona{
		nombre: "Nicolas",
	}
	DiAlgo(&p1)
}