package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Print("Enter a multiplier: ")
	var multiplier float64
	fmt.Scanf("%f", &multiplier)
	output := input * multiplier
	fmt.Println(output)
}
