package main

import (
	"encoding/json"
	f "fmt"
	"os"
	
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

func main() {

	ul := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Dale boca",
			"Que hace fabra",
			"Como va a jugar de titular ramirez",
		},
	}

	ul1 := usuario{
		Nombre:   "Nicolas",
		Apellido: "Gonzale",
		Edad:     78,
		Dichos: []string{
			"Ahboe",
			"Que digo",
			"falso",
		},
	}

	usuarios := []usuario{ul,ul1}

	// f.Println( usuarios )

	err := json.NewEncoder( os.Stdout ).Encode(usuarios)

	if err != nil{
		f.Println(err)
	}

	f.Println( usuarios )

}
