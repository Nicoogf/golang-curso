package main

import (
	f "fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	f.Println("Numero de CPUs : ", runtime.NumCPU())
	f.Println("Numero de Goroutines : ", runtime.NumGoroutine())
	var contador int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			f.Println("Contador : ", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		f.Println("Numero de Goroutines : ", runtime.NumGoroutine())
	}

	wg.Wait()
	f.Println("Cuenta : ", contador)

}
