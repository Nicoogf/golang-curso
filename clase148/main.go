package main

import (
	f "fmt"
	"runtime"
	// "sync"
)

func main() {
	f.Printf("Numeros de CPU : %v \n" , runtime.NumCPU())
	f.Printf("Numeros de goroutines : %v \n" , runtime.NumGoroutine())

	// var wg sync.WaitGroup
	
	go func(){
		f.Println("Hola desde la primer goroutine")
	}()

	go func(){
		f.Println("Hola desde la segunda goroutine")
	}()

	f.Printf("Numeros de CPU : %v \n" , runtime.NumCPU())
	f.Printf("Numeros de goroutines : %v \n" , runtime.NumGoroutine())


	f.Println("A punto de finalizar main")

	f.Printf("Numeros de CPU : %v \n" , runtime.NumCPU())
	f.Printf("Numeros de goroutines : %v \n" , runtime.NumGoroutine())
}