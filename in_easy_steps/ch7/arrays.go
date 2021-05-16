package main

import (
	"fmt"
)

func main() {
	var cars [4]string
	cars[0] = "honda civic"
	cars[1] = "ford mustang"
	cars[2] = "chevy volt"
	cars[3] = "toyota corolla"

	coords := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Cars:", cars)
	fmt.Println("Coords:", coords)
	fmt.Println()
	fmt.Println()
	for i := range cars {
		fmt.Print("[",cars[i], "]")
		fmt.Print(",")
	}
	fmt.Println()
	fmt.Println()
	for coord := range coords {
		fmt.Println(coords[coord])
	}
}
/* Cars: [honda civic ford mustang chevy volt toyota corolla]
Coords: [[1 2 3] [4 5 6]]


[honda civic],[ford mustang],[chevy volt],[toyota corolla],

[1 2 3]
[4 5 6]
*/
