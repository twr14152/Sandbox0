//https://play.golang.org/p/jFvmuVci5T5
package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.Pi
	fmt.Println("Pi:", pi)
	fmt.Println("\nFloor:", math.Floor(pi))
	fmt.Println("\nCeiling:", math.Ceil(pi))
	fmt.Println("\nRound:", math.Round(pi))
	fmt.Println("\nTruncated:", math.Trunc(pi))
}

/*
Pi: 3.141592653589793

Floor: 3

Ceiling: 4

Round: 3

Truncated: 3
*/
