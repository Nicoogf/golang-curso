package main

import (
	f "fmt"
)

func main() {
	foo()
	func ()  {
		f.Println("Ejecucion de funcion anonima")
	}()
}


func foo () {
	f.Println("Ejecucion de la funcion Foo")
}