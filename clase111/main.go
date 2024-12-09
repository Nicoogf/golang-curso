package main

import (
	f "fmt"
)

func main(){
	g := foo()
	f.Println( g() )
	f.Println( g() )
	f.Println( g() )

	h := foo()
	f.Println( h() )
	f.Println( h() )
	f.Println( h() )


}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}