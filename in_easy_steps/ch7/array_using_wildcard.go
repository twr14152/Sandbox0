package main

import (
	"fmt"
)

func main() {
	// Wildcard [...] using literal values
	cars := [...]string{"honda civic", "ford mustang", "chevy volt", "toyota corolla"}

	fmt.Println("Cars:", cars)

	fmt.Println()
	fmt.Println("Using for loop to range through the array elements:")
	for i := range cars {
		fmt.Printf("index: %v value:[%v]\n", i, cars[i])
	}
	fmt.Println()
	fmt.Println()
}

/*
Cars: [honda civic ford mustang chevy volt toyota corolla]

Using for loop to range through the array elements:
index: 0 value:[honda civic]
index: 1 value:[ford mustang]
index: 2 value:[chevy volt]
index: 3 value:[toyota corolla]

*/
