package main

import (
	f "fmt"
)

func main () {
	x := 0
	f.Println("El valor de x antes " ,  x)
	f.Println("El valor de x antes " ,  &x)
	foo(&x)
	f.Println("Resultado final :" , x)
	
}

func foo ( y *int ) {
	f.Println( "El valor antiguo" , y )
	f.Println( "en direccion" , *y )
	*y = 42
	f.Println( "El valor despues" , y )
	f.Println( "en direccion despues" , *y )

}