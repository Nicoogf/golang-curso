package main

import (
	"fmt"
)


// func main() {
// 	ca :=  make( chan int , 2 )

// 	ca <- 78
// 	ca <- 79


// 	fmt.Println(<- ca )
// 	fmt.Println(<- ca )
// 	fmt.Println("-----------")
// 	fmt.Printf("%T \n" , ca )
// }


func main() {

	ca := make(<-chan int , 2 )

	ca <- 42
	ca <- 43

	fmt.Println(<- ca )

}