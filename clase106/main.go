package main

import (
	f "fmt"
	m "math"
)


type circulo struct{
	radio float64
}

type cuadrado struct{
	longitud float64
}


func ( c circulo ) area () float64 {
	return  m.Pi * (c.radio * c.radio)
}

func ( c cuadrado ) area () float64 {
	return c.longitud * c.longitud
}

type forma interface {
	area() float64
}

func info ( forma forma ) {
	f.Println( forma.area() )
}

func main() {
	circ := circulo{radio: 12.345,}
	cuadrado := cuadrado{longitud: 15,}

	info(circ)
	info(cuadrado)

}