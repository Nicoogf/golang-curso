package main

import (
	"encoding/json"
	_ "encoding/json"
	f "fmt"
)

type Persona struct{
	Nombre string
	Apellido string
	Edad int
}

func main () {
	s := `[{"Nombre":"Nicolas","Apellido":"Gonzales","Edad":28},{"Nombre":"Adrian","Apellido":"Nario","Edad":48}]`
	bs := []byte(s)

	f.Printf("%T \n" ,  s)
	f.Printf("%T \n" ,  bs)

	var personas []Persona

	err := json.Unmarshal( bs , &personas)

	if err != nil {
		f.Println(err)
	}

	f.Println("Toda la data : " , personas)

	for i , v := range personas{
		f.Println("Indice " , i  , " Usuario " , v.Nombre)
	}



}