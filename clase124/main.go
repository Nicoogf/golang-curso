package main

import (
	f "fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "clave123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
		if err != nil {
			f.Println(err)
		}
	hash := string(bs)
	f.Println(s)
	f.Println(bs)
	f.Println(string(hash))

	claveLogin := "clave123"
	err = bcrypt.CompareHashAndPassword(bs , []byte(claveLogin))
	if err != nil {
		f.Println("No se puede loguear")
		return
	}
	f.Println("Te haz logueado")
}
