package main

import "fmt"

func deferingAction(){
    fmt.Println("Hello World")
    defer fmt.Println("First Deferred Print Line")
    fmt.Println("This will print before the first defer")
    defer fmt.Println("Second Deffered Print Line")
    fmt.Println("Finally, we'll leave the function.")
}

func main(){
    deferingAction()
}
