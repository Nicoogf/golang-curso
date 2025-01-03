package main

import(
	"fmt"
	_"runtime"
)

func main(){
	ca :=  make( chan int )

	go enviar(ca)

	go recibir(ca)

	fmt.Println("Finalizado.")

}


func enviar(c chan<- int){
	c <- 42
}

func recibir(c <-chan int){
	fmt.Println(<- c)
}