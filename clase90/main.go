package main

import (
	f "fmt"
)

func main() {
	foo()
	bar("Carlos")

	s1 := woo("Pope")
	f.Println(s1)

	x , y := saludar("Nacho" , "Pibe")
	f.Println(x)
	f.Println(y)
}

func foo () {
	f.Println("Prueba de funcion")
}

func bar ( s string ){
	f.Println("Saludo a : "  , s)
}

func woo (st string) string {
	return f.Sprint("Hola desde funcion woo saludo a " , st)
}

func saludar ( n , a string ) (string , bool) {
	p := f.Sprint( n , " " , a , ` dice "hola" `)
	q :=  true
	return p,q
}