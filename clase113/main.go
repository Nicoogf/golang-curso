package main

import (
	f "fmt"
)

func main() {
	a := 42

	f.Println(&a)
	f.Printf("%T \n" , a)
	f.Printf("%T \n" , &a)

	var b *int = &a
	f.Println("El valor de b solo : " , b)
	f.Println("El valor de b con asterisco : " , *b)
	f.Println("El valor de b con asterisco : " , &b)

	*b = 44

	f.Println("Ahora A vale : " ,  a)

}