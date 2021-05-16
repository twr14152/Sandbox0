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
	fmt.Println("Using for loop to range through the array elements:")
	for i := range cars {
		fmt.Println(i,"[",cars[i],"]")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("For loop through array using len(array):")
	for i := 0; i < len(cars); i++ {
		fmt.Print("[", cars[i], "]")
		fmt.Print(",")
	}
	fmt.Println()
	fmt.Println()
	for coord := range coords {
		fmt.Println(coords[coord])
	}
}

/* 
Cars: [honda civic ford mustang chevy volt toyota corolla]
Coords: [[1 2 3] [4 5 6]]

Using for loop to range through the array elements:
0 [ honda civic ]
1 [ ford mustang ]
2 [ chevy volt ]
3 [ toyota corolla ]


For loop through array using len(array):
[honda civic],[ford mustang],[chevy volt],[toyota corolla],

[1 2 3]
[4 5 6]
*/
