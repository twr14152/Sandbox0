package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Outer loop number ", i)
		for x := 1; x <= 3; x++ {
			fmt.Println("\tInner loop number ", x)
		}
	}
}

/*
Outer loop number  1
	Inner loop number  1
	Inner loop number  2
	Inner loop number  3
Outer loop number  2
	Inner loop number  1
	Inner loop number  2
	Inner loop number  3
Outer loop number  3
	Inner loop number  1
	Inner loop number  2
	Inner loop number  3
*/
