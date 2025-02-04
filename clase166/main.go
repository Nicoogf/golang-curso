package main

import (
	f "fmt"
)

func main(){
	c := gen()
	recibir(c)
	f.Println("A punto de finalizar")
}

func recibir( c <- chan int){
	for v := range c {
		f.Println(v)
	}

}

func gen() <- chan int {
	c := make(chan int)

	go func(){
		for i := 0 ; i < 10 ; i++ {
			c <- i
		}
		close(c)
	}()
	

	return c
}