package main

import (
	"fmt"
)

func main() {
	//unbuffer o sin buffer
	ca := make(chan int)

	go func(){
		ca <- 45
	}()

	fmt.Println( <- ca  )


}
