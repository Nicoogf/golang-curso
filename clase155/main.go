package main

import (
	"fmt"
)

// func main() {
// 	//unbuffer o sin buffer
// 	ca := make(chan int)

// 	go func(){
// 		ca <- 45
// 	}()

// 	fmt.Println( <- ca  )


// }


// func main(){
// 	ca := make( chan int , 1 )
// 	ca <- 42

// 	fmt.Println( <- ca )
// }


func main(){
	ca := make(chan int , 2 )
	ca <- 42
	ca <- 43

	fmt.Println( <- ca)
	fmt.Println( <- ca)
}