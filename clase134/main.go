package main

import (
	f "fmt"
	"runtime"
	"sync"
	
)

func main() {
	f.Println("Numero de CPUs : " , runtime.NumCPU() )
	f.Println("Numero de Goroutines : " , runtime.NumGoroutine() )
	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0 ; i < gs ; i++ {
		go func() {
			v := contador
			v++		
			runtime.Gosched()
			contador = v
			wg.Done()
		}()
		f.Println("Numero de Goroutines : " , runtime.NumGoroutine() )
	}

	wg.Wait()
	f.Println("Cuenta : " , contador )

}