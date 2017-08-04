package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var val1 float64
	fmt.Scanf("%f", &val1)
	fmt.Print("Enter a second number: ")
	var val2 float64
	fmt.Scanf("%f", &val2)
	add := val1 + val2
	// fmt.Printf will print values of %+v Println will not
	fmt.Printf("%+v + %+v = %+v\n", val1, val2, add)
	sub := val1 - val2
	fmt.Printf("%+v - %+v = %+v\n", val1, val2, sub)
	mult := val1 * val2
	fmt.Printf("%+v * %+v = %+v\n", val1, val2, mult)
	div := val1 / val2
	fmt.Printf("%+v / %+v = %+v\n", val1, val2, div)

}
