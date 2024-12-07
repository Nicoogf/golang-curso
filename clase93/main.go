package main

import (
	f "fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	f.Println("Foo")
}

func bar() {
	f.Println("Bar")
}
