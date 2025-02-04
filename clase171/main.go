package main

import (
	f "fmt"
)

func main(){
	var respuesta1, respuesta2, respuesta3 string 

	f.Print("Nombre : ")
	_, err := f.Scan(&respuesta1)
	if err != nil{
		panic(err)
	}

	f.Print("Comida Favorita : ")
	_, err = f.Scan(&respuesta2)
	if err != nil{
		panic(err)
	}

	f.Print("Deporte favorito : ")
	_, err = f.Scan(&respuesta3)
	if err != nil{
		panic(err)
	}

	f.Println(respuesta1,respuesta2,respuesta3)
}