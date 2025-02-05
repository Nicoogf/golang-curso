package main

import (
	"errors"
	"fmt"

	"log"
)

var errorMath =  errors.New("no puede ser raiz cuadrada real de un numero negativo")

func main(){
	fmt.Printf("%T\n" , errorMath)
	_ , err := sqrt(-10)
	if err != nil{
		log.Fatalln(err)
	}
}


func sqrt( f float64) (float64 , error ){
	if f < 0 {
		return 0 , errorMath
	}
	return 42 , nil
}