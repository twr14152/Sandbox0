package main

import (
	"fmt"
)

func main() {
	if 'A' == 'A' {
		fmt.Println("\nCharacters Match")
	}
	if 5 > 1 {
		if 7 > 2 {
			fmt.Println("\nBoth expressions are true")
		}
	}
	if 5 < 1 {
		fmt.Println("\n1st Condition is true")
	} else if 'A' != 'A' {
		fmt.Println("\n2nd Condition is true")
	} else {
		fmt.Println("\nBoth Conditions are false")
	}
}

/*
Characters Match

Both expressions are true

Both Conditions are false
*/
