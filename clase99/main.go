package main

import (
	f "fmt"
)

func main() {
	li := []int {1,2,3,4,5,6,7,8,9}
	s:= suma(li...)
	f.Println(s)

	s1 := sumaPares( suma , li...)
	f.Println(s1)

	
}

func suma( xi ...int) int {
		 f.Printf("%T \n" ,xi)
		 total := 0
		 for _ , v := range xi{
			total += v
		 }
		 return total  
}

func sumaPares (f func(xi ...int) int , vi ...int) int {
	var y []int
	for _ , v := range vi{
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	return f ( y... )
} 