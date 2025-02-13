package main

import (
	"fmt"
	_ "fmt"
	"log"
	"os"
)

func main() {
	// _ , err := os.Open("sin-archivo.txt")
	// if err != nil {
	// 	fmt.Println("Ocurrio un error : " , err)
	// }

	// _ , err := os.Open("sin-archivo.txt")
	// if err != nil {
	// 	log.Println("Ocurrio un error : " , err)
	// }

	f , err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2 , err := os.Open("sin-archivo.txt")

	if err != nil {
		log.Println("Ocurrio un error : " , err)
	}

	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio")
}