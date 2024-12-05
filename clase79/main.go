package main

import(
	f "fmt"
)

func main() {
	mapa_nuevo :=  map[string][]string{
		"Colores" : {"Blanco", "Amarrillo"},
	}


	mapa_nuevo["Numeros"] = []string{"Uno","Dos"}

	for k, v := range mapa_nuevo{
		f.Println(k)
		for _ , j := range v {
			f.Println(j)
		}
	}

	
}