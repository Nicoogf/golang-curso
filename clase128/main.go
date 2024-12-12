package main

import (
	f "fmt"
	"sort"
)

func main(){
	si := []int {8,9,12,45,78,95,64,75}
	f.Println(si)
	sort.Ints(si)
	f.Println(si)

	ss := []string{"Hola","Chau","Dice","Boca","Copa","Libertadores"}
	f.Println(ss)
	sort.Strings(ss)
	f.Println(ss)


}