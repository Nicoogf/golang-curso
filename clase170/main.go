package main

import (
	f "fmt"
)

func main(){
	c := make(chan int)

	for j := 0 ;j < 10 ; j++ {
		go func (){
			for i := 0 ; i<100 ; i++{
				c <- i
			}			
		}()
	}

	for v := range c {
		f.Println(v)
	}

	f.Println("a punto de finalizar")
}