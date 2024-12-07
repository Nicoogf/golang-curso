package main

import(
	f "fmt"
)

type persona struct{
	nombre string
	apellido string
}

type agenteSecreto struct{
	persona
	lpm bool
}

func ( s agenteSecreto) presentar() {
	f.Println("Hola soy" , s.nombre , " " , s.apellido , "el agente secreto")
}


func ( p persona) presentar() {
	f.Println("Hola soy" , p.nombre , " " , p.apellido , " persona")
}


type humano interface{
	presentar()
}

func bar ( h humano ){
	switch h.(type){
		case persona:
			f.Println("Fui pasado a la funcion bar (Persona)" , h.(persona).nombre)
		case agenteSecreto:
			f.Println("Fui pasado a la funcion bar (agente secreto)" , h.(agenteSecreto).nombre)
	}
	f.Println("Fui pasado a la funcion bar" , h)
}

type manzana int 

func main() {

	as1 := agenteSecreto{
		persona: persona{
			nombre: "Eduar",
			apellido: "Tua",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre: "Miss",
			apellido: "MoneyPenny",
		},
		lpm: true,
	}

	p := persona{
		nombre: "Carito",
		apellido: "Guz",
	}

	f.Println(as1)
	
	as1.presentar()
	as2.presentar()

	bar(as1)
	bar(as2)
	bar(p)

	var x manzana  = 42

	f.Println(x) 
	f.Printf("%T \n" , x)
	var y int
	y = int(x)
	f.Printf("%T" , y)
}
