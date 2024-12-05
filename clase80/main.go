package main

import (
	f "fmt"
)


func main(){
	mapa_ochenta :=  map[string][]string {
		"Juan" : {"Azul", "Blanco" , "Rojo" },
		"Jose" : {"Verde", "Amarillo" , "Negro" },
	}


	mapa_ochenta["Adrian"] = []string{
		"Naranja" , "Gris" , "Celeste" ,
	}


	for k , v := range mapa_ochenta {
		f.Println("El usuario " , k)
		for i := 0 ; i < len(v) ; i ++ {
			f.Println("Color " , i+1 , " fue " , v[i])
		}
	}
}