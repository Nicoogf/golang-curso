package main

import(
	f "fmt"
)

func main() {
	c := make(chan int)

	go enviar(c)

	recibir(c)

	f.Println("Finalizado")
}

func enviar(c chan<-int ){
	c <- 42
}

func recibir(c <-chan int ){
	f.Println(<-c)
}


