//https://play.golang.org/p/fJi3iY8Vf14
package main

import (
	"fmt"
	"math"
)

func main() {
	var rad, area, perim float64
	rad = 4
	fmt.Printf("Radius if a Circle: %.2f\n", rad)
	area = math.Pi * (rad * rad)
	fmt.Printf("\nArea of a Circle: %.2f \n", area)
	perim = 2 * (math.Pi * rad)
	fmt.Printf("\nPerimeter of a Circle: %.2f\n", perim)

}
/*
Radius if a Circle: 4.00

Area of a Circle: 50.27 

Perimeter of a Circle: 25.13
*/
