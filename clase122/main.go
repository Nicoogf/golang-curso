package main

import (
	f "fmt"
	"sort"
)

func main() {

	xi := []int {2 ,4,8,9,6}
	xs := []string {"a","b", "v" , "p" , "k"}

	sort.Ints(xi)
	sort.Strings(xs)
	f.Println(xi)
	f.Println(xs)

}