package main
import(
	f "fmt"
)

func main() {
	c := make(chan int)

	go func (){
		c <- 42
	}()

	v , ok := <- c
	f.Println(v,ok)
	
	close(c)

	v ,ok = <- c
	f.Println(v,ok)
}