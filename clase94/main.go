package main

import(
	f "fmt"
)

type mascota struct{
	nombre string
	nada bool
}

type perro struct{
	mascota
	raza string
}

func ( identificador perro ) ruido() string{
	datoRetornable := f.Sprintln("El perro " , identificador.nombre , "hizo guau")
	return datoRetornable
}

func main(){

	dola := perro{
		mascota: mascota{
			nombre: "Carla",
			nada: false,
		},
		raza: "Labradora",
	}

	f.Println(dola.ruido())

}