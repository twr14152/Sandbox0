package main

import (
	"fmt"
)

var a int

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)

	}
}
func main() {
	for i := 0; i <= 40; i++ {
		//fmt.Println("fibanacci for even and odd numbers ", fibonacci(i), " ")
		if fibonacci(i)%2 == 0 && fibonacci(i) < 4000000 {
			//fmt.Println("fibanacci for even value of i is ", i)
			fmt.Println("even values", fibonacci(i))
			a = fibonacci(i) + a
			fmt.Println("sum of a is ", a)
		} else {
			continue
			//fmt.Print("fibanacci for odd value of i is ", i)
			//fmt.Print(fibonacci(i), " ")
		}
	}
}
