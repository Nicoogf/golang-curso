package main

import (
	f "fmt"
)

func main() {
	a:= incrementor()
	b := incrementor()
	f.Println(a())
	f.Println(a())
	f.Println(a())
	f.Println(a())
	f.Println(b())
	f.Println(b())


}

func incrementor() func() int {
	var x int
	return func () int {
		x++
		return x
	}
}