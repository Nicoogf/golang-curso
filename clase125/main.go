package main

import(
	"fmt"
	"encoding/json"
)

type Persona struct{
	Nombre string
	Edad int
}

func main(){

	p1 :=  Persona{Nombre: "Nicolas" , Edad: 28}
	p2 :=  Persona{Nombre: "Gabriel" , Edad: 27}
	p3 :=  Persona{Nombre: "Juan" , Edad: 25}

	usuarios := []Persona{p1,p2,p3}
	fmt.Println(usuarios)

	bs , err := json.Marshal(usuarios)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("El resultado fue :" , string(bs))


}