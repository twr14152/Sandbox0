//https://play.golang.org/p/B2-0ynDEILR
package main

import (
	"fmt"
	"math"
)

func main() {
	square := math.Pow(5, 2)
	cube := math.Pow(4, 3)
	fmt.Println("\nLargest Positive:", math.Max(square, cube))
	fmt.Println("\nSmallest Positive:", math.Min(square, cube))

	square *= -1
	cube *= -1

	fmt.Println("\nLargest Negative:", math.Max(square, cube))
	fmt.Println("\nSmallest Negative:", math.Min(square, cube))
}

/*
Largest Positive: 64

Smallest Positive: 25

Largest Negative: -25

Smallest Negative: -64

*/
