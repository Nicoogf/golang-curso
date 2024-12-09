package main

import (
	f "fmt"
)

func main() {
	g := func( xi []int )int{
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi) - 1 ]
	}

	x := foo( g , []int {1,2,3,4,5})
	f.Println(x)
}

func foo( ejemplo func([]int) int , arreglo []int ) int {
	n := ejemplo( arreglo ) 
	n++
	return n
}