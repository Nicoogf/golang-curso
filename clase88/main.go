package main

import (
	f "fmt"
)

type vehiculo struct{
	puertas int
	color string
}

type camion struct{
	vehiculo
	cuatroRuedas bool
}

type sedan struct{
	vehiculo
	lujoso bool
}

func main() {

	camion_uno := camion{
	 vehiculo: vehiculo{puertas: 4,color: "Gris"},
	 cuatroRuedas: true,
	}
	sedan_uno := sedan{
		vehiculo: vehiculo{puertas: 4 , color: "Rojo"},
		lujoso: false,
	}

	f.Println("Informacion de Camion")
	f.Println("Puertas : " ,  camion_uno.puertas)
	f.Println("Color : " ,  camion_uno.color)
	f.Println("Cuatro ruedas : " ,  camion_uno.cuatroRuedas)

	f.Println("-----------------------")

	f.Println("Informacion de Sedan")
	f.Println("Puertas : " ,  sedan_uno.puertas)
	f.Println("Color : " ,  sedan_uno.color)
	f.Println("Cuatro ruedas : " ,  sedan_uno.lujoso)



}