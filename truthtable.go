package main

import "fmt"

func main() {
	//logical anding
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", true && false)
	//logical oring
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false || false)
	//!x = Notx
	fmt.Println("!false = ", !false)
	fmt.Println("!true = ", !true)
}
