package main

import (
	f "fmt"
)

func main(){
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar
	go enviar( par , impar , salir)

	//recibir
	recibir(par ,impar ,salir)

	f.Println("Finalizar")

}


//Enviar

func enviar( p , i , s chan <- int) {
	for j := 0 ; j < 100 ; j++ {
		if j % 2 == 0{
			p <- j
		}else{
			i <- j
		}
	}
	s <- 0
}

func recibir( p , i ,  s <- chan int){
	for{
		select{
		case v:= <-p :
			f.Println("Desde el canal par : " , v)		
		case v:= <-i :
		f.Println("Desde el canal impar : " , v)
	
		case v:= <-s :
		f.Println("Desde el canal salir : " , v)
		return
		}
	}
}

