package main

import(
	"fmt"
	"sort"
)


type Persona struct{
	Nombre string
	Edad int
}

type PorNombre []Persona

func ( a PorNombre ) Len() int {
	return len(a)
}

func(a PorNombre) Swap ( i , j  int ) {
	a[i] ,  a[j] = a[j] , a[i]
}

func(a PorNombre) Less ( i , j  int ) bool {
	return a[i].Nombre < a[j].Nombre 
}


func main () {
	p1 := Persona {"Eduardo" , 32 }
	p2 := Persona {"Maria" , 25 }
	p3 := Persona {"Carolina" , 56 }
	p4 := Persona {"Adriana" , 60 }

	personas := []Persona{p1,p2,p3,p4}

	fmt.Println(personas)

	sort.Sort(PorNombre(personas))

	fmt.Println(personas)
	
	
	
}