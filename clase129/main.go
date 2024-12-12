package main

import (
	f "fmt"
	"sort"
)

type usuario struct{
	Nombre string
	Apellido string
	Edad  int
	Dichos []string
}

type PorEdad []usuario

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }


type PorApellido []usuario

func (a PorApellido) Len() int           { return len(a) }
func (a PorApellido) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }



func main (){
	user_uno := usuario{Nombre: "Nicolas",Apellido: "Gonzales",Edad: 28,Dichos: []string{"Como que juega fabra","Ramirez se tiene que ir","Como lo van a empatar a ultimo momento"}}
	user_dos := usuario{Nombre: "Gabriel",Apellido: "Palacios",Edad: 48,Dichos: []string{"Que numero dijiste","La defenza no para a uno","este año es"}}
	user_tres := usuario{Nombre: "Desconocido",Apellido: "NN",Edad: 10,Dichos: []string{"No dije nada"}}

	usuarios := []usuario{user_uno,user_dos,user_tres}

	for _ , valor := range usuarios {
		f.Println(valor.Nombre , valor.Apellido , valor.Edad)
		for _ , valor := range valor.Dichos{
			f.Println("\t - " ,valor)
		}
		
	}

	sort.Sort(PorEdad(usuarios))
	f.Println(usuarios)

	for _ , valor := range usuarios {
		f.Println(valor.Nombre , valor.Apellido , valor.Edad)

		sort.Strings(valor.Dichos)
		for _ , valor := range valor.Dichos{
			f.Println("\t - " ,valor)
		}
		
	}

	sort.Sort(PorApellido(usuarios))
	f.Println(usuarios)

	for _ , valor := range usuarios {
		f.Println(valor.Nombre , valor.Apellido , valor.Edad)
		for _ , valor := range valor.Dichos{
			f.Println("\t - " ,valor)
		}
		
	}


}