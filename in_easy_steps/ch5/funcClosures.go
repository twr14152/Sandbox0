//First Class Function - allows functions to be assigned to variables
//Anonymous function - function definition with out having a name
//Closures - a function nested inside an outer function that retains
//access to variables declared in outer function.

package main

import (
	"fmt"
)

func main() {
	area := func(length, width int) int {
		return length * width
	}
	fmt.Printf("area Type: %T\n", area)
	fmt.Println("Area 1:", area(10, 4))
	fmt.Println("Area 2:", area(12, 5))
	counter := func() func() int {
		num := 0
		return func() int {
			num++
			return num
		}
	}()
	fmt.Printf("counter type: %T\n", counter)
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())

}

/*
area Type: func(int, int) int
Area 1: 40
Area 2: 60
counter type: func() int
Count: 1
Count: 2
Count: 3
*/
