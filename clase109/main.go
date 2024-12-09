package main

import (
	f "fmt"
)

func main() {
	fun := foo()
	f.Println(fun())
}

func foo() func() int {
	return func() int{
		return 45
	}
}