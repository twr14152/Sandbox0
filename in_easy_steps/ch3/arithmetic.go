package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 4
	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Remainder:\t", (a % b))
	a++
	fmt.Println("Increment:\t", a)
	b--
	fmt.Println("Decrementer:\t", b)
	
	
}
/*
Addition:	 9
Subtraction:	 1
Multiplication:	 20
Division:	 1
Remainder:	 1
Increment:	 6
Decrementer:	 3
*/
