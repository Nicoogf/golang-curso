package main

import (
	f "fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main(){
	f.Println("OS \t " , runtime.GOOS )
	f.Println("ARCH \t " , runtime.GOARCH )
	f.Println("CPU \t " , runtime.NumCPU() )
	f.Println("Goroutines \t " , runtime.NumGoroutine() )

	wg.Add(1)
	go foo()
	bar()

	f.Println("CPU \t " , runtime.NumCPU() )
	f.Println("Goroutines \t " , runtime.NumGoroutine() )
	

	wg.Wait()
}

func foo (){
	for i:=0 ; i<5 ; i++{
		f.Println("foo - " , i)
	}
	wg.Done()
}

func bar (){
	for i:=0 ; i<5 ; i++{
		f.Println("bar - " , i)
	}
}