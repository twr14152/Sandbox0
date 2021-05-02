package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159
	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)
	fmt.Printf("Pi approximately: %v \n\n", pi)
	fmt.Printf("Red: %v point \n", red)
	fmt.Printf("Blue: %v point \n", blue)
	fmt.Printf("Black: %v points \n", black)
	fmt.Printf("Pink = %v, (%v * 2) = %v", pink, pink, pink * 2)
}
/*
Pi approximately: 3.14159 

Red: 1 point 
Blue: 5 point 
Black: 7 points 
Pink = 6, (6 * 2) = 12
*/
