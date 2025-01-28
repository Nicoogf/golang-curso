// package main

// import (
// 	f "fmt"
// )

// func main(){
// 	par := make(chan int)
// 	impar := make(chan int)
// 	salir := make(chan int)

// 	go enviar( par , impar, salir)

// 	recibir( par, impar , salir)

// 	f.Println("Finalizado")

// }

// func enviar( par , impar , salir chan <- int ){
// 	for i := 0 ; i < 100 ; i++ {
// 		if i % 2 == 0{
// 			par <- i
// 		}else{
// 			impar <- i
// 		}
// 	}
// 	close(salir)
// }

// func recibir( par , impar  , salir <-chan int ) {
// 	for {
// 		select {
// 			case v := <- par:
// 			f.Println("El valor recibido del canal par : " , v)
// 			case v := <- impar:
// 			f.Println("El valor recibido del canarl impar  : " , v)
// 			case i , ok := <-salir :
// 				if !ok {
// 					f.Println("Desde coma ok " , i , ok)
// 					return
// 				}else{
// 					f.Println("Desde coma ok " , i , ok)
// 				}
// 		}
// 	}
// }


package main

import (
	f "fmt"
)

func main() {
	c := make( chan int)

	go func () {
		c <-42
		close(c)
	}()

	v , ok := <- c

	f.Println(v, ok)

	
	v , ok = <- c

	f.Println(v, ok)
}