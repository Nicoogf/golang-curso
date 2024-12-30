package main 

import (
	f "fmt"
	"runtime"
)


func main() {
	numero := 45 
	f.Println(numero)

	f.Printf("El SO es : %s , con arquitectura : %s " , runtime.GOOS , runtime.GOARCH)
}