package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f , err := os.Create("nombres.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Nicolas")

	io.Copy(f , r)
	test()
}


func test() {
	f, err := os.Open("nombres.txt")
	if err != nil {
		fmt.Println(err)
		return
	} 

	defer f.Close()

	bs , err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}