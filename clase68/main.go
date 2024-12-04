package main

import(
	f "fmt"
)

func main(){
	m := map[string] int {
		"Batman" : 32 ,
		"Robin" : 27,
	}

	f.Println(m)
	f.Println(m["Batman"])
	f.Println(m["Robin"])
	f.Println(m["Barco"])
	
	v , ok :=  m["Eduardo"]
	b , okdos :=  m["Batman"]
	f.Println(v , ok)
	f.Println(b , okdos)

	// if v,ok := m["Batman"] ; ok {
	// 	f.Println("Imprimio el ok" , ok , v)
	// }

	// if v,ok := m["Eduardo"] ; ok {
	// 	f.Println("Imprimio el ok" , ok , v)
	// }

	// m["Eduardo"] = 25
	// f.Println(m["Eduardo"])
	// f.Println(m)

	// for llave , valor := range m {
	// 	f.Println( llave , valor )
	// }

	// for clave , valor :=  range m {
	// 	f.Println(clave , valor)
	// }

	for clave , valor := range m {
		f.Println(clave , valor)
	}
}