package main

import (
	f "fmt"
)

func main() {
	arreglo := []int {4,9}
	resultado := foo(arreglo...)
	resultado_dos := bar(arreglo)

	f.Println(resultado)
	f.Println(resultado_dos)

}

func foo( n ...int ) int {
	total := 0
	for _ , v :=  range n {
       total += v
	}
	return total
}

func bar ( n []int ) int {
	total := 0
	for _ , v := range n{
		total += v
	}
	return total
}