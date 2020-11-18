package main

import (
	"fmt"
)

func main() {
	days := [7]string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	weekend := days[5:]

	fmt.Printf("Slice weekend: %v\n", weekend)
	fmt.Printf("Type weekend: %T\n", weekend)
	fmt.Printf("Length weekend: %v \n", len(weekend))
	fmt.Printf("Capacity weekend: %v \n", cap(weekend))
	fmt.Printf("Pointer weekend: %p \n", weekend)
	fmt.Printf("Address days[0]: %p \n", &days[5])

}

/*
Slice weekend: [sat sun]
Type weekend: []string
Length weekend: 2 
Capacity weekend: 2 
Pointer weekend: 0xc00007c050 
Address days[0]: 0xc00007c050 
*/
