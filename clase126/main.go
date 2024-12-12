package main


import (
	f "fmt"
	"encoding/json"
)

type Persona struct{
	Nombre string
	Edad int
}

func main() {
	codigo := `[{"Nombre":"Nicolas","Edad":28},{"Nombre":"Gabriel","Edad":27},{"Nombre":"Juan","Edad":25}]`
	bitesString := []byte(codigo)

	var personas []Persona 

	err := json.Unmarshal(bitesString , &personas)

	if err != nil {
		f.Println(err)
	}

	f.Println("La data parseada : " , personas)
	f.Printf("El tipo de dato es : %T" , personas)

	for i , p :=  range personas{
		f.Printf("En el indice %d , esta el usuario %s \n" , i , p.Nombre)
	}



}