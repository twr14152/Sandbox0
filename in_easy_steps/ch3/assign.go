package main

import (
	"fmt"
)

func main() {
	a, b := 8, 4
	fmt.Println("Assigned Value:\ta =", a, "\tb = ", b)
	a += b
	fmt.Println("Add and Assign:\ta =", a)
	a -= b 
	fmt.Println("Subtract and Assign:\ta =", a)
	a *= b
	fmt.Println("Multiply and Assign;\ta =", a)
	a /= b
	fmt.Println("Divide and Assign:\ta =", a)
	a %= b
	fmt.Println("Remainder Assign:\ta =", a)
}
/*
Assigned Value:	a = 8 	b =  4
Add and Assign:	a = 12
Subtract and Assign:	a = 8
Multiply and Assign;	a = 32
Divide and Assign:	a = 8
Remainder Assign:	a = 0

*/
