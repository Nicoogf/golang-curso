package main

import(
	f "fmt"
)

func main() {
	mapa_uno := map[string][]string {
		"Eduardo" :  {"Computadora" , "montaña" , "playa"} ,
		"Carlos" :  {"Leer" , "comprar" , "musica"} ,
		"Juan" :  {"Helado" , "pinto" , "bailar"} ,
	}


	for k , v := range mapa_uno{
		f.Println("Clave : " , k , "Valor :"  , v)
	}


}