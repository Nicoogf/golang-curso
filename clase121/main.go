package main

import(
    i "io"
    f "fmt"
    "os"
)

func main(){
    f.Println("Hola mundo")
    i.WriteString(os.Stdout , "Hola mundo")
    i.WriteString(os.Stdout , "Hola mundo")


}