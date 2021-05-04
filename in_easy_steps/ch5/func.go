package main

import (
	"fmt"
)
func first() {
	msg := "Hello from first function"
	fmt.Println(msg)
}

func sqFive() {
	fmt.Printf("%v\n", 5*5)
}

func main() {
	first()
	fmt.Print("5x5=")
	sqFive()
}

/*
Hello from first function
5x5=25
*/
