package main

import (
	"fmt"
	"math" // used to determine how close to answer
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * x)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("My programs best guess is : ", Sqrt(2))
	fmt.Println("The answer is: ", math.Sqrt(2))
}

/*
1.25
1.359375
1.39739990234375
1.4092182805761695
1.4127442399986556
1.4137826680863108
1.414087309940999
1.414176579906956
1.4142027301176223
1.4142103896495881
My programs best guess is :  1.4142103896495881
The answer is:  1.4142135623730951

Program exited.
*/
