package main

import (
	f "fmt"
)

func main() {
	 n1 := foo(4)
	 f.Println("La funcion foo retorno un int " , n1)

	 n2,n3 := bar(10 , "Hola mundo")
	 f.Println("El entero de la funcion bar fue" , n2 )
	 f.Println("El string de la funcion bar fue ", n3)
}

func foo( n int ) int {
	return n
}

func bar(n int , y string) (int , string) {
	return n , y 
}