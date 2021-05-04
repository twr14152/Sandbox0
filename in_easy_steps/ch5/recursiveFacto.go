//Recursive func

package main

import (
	"fmt"
)

func facto(num int) int {
	if num == 0 {
		return 1
	}
	return num * facto(num-1)
}

func main() {
	for i := 1; i <= 7; i++ {
		fmt.Println("Factorial", i, "=", facto(i))
	}

}

/*
Factorial 1 = 1
Factorial 2 = 2
Factorial 3 = 6
Factorial 4 = 24
Factorial 5 = 120
Factorial 6 = 720
Factorial 7 = 5040
*/
