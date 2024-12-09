package main

import (
	f "fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func () {
		f.Println("Ejecucion interna de foo")
	}()
	f.Println("Ejecucion de Foo")
}

func bar() {
	f.Println("Ejecucion de Bar")
}