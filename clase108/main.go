package main

import(
	f "fmt"
)

var x int 
var g func() = func() {
	f.Println("Imprimo g desde afuera")
}

func main(){
	fun := func(){
		for i := 0 ; i < 5 ; i++ {
			f.Println("Valor " , i)
		}
	}

	fun()
	f.Printf("%T \n" , fun)
	f.Printf("%T \n" , x)
	f.Printf("%T \n" , g)
	f.Println("Listo")

	g()

}