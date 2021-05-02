package main

import (
	"fmt"
)

func main() {
	var yes, no bool = true, false
	result := (yes && no)
	fmt.Println("AND Logic:\tyes && no\t", result)
	result = (yes || no)
	fmt.Println("OR Logic:\tyes || no\t", result)
	result = !yes
	fmt.Println("NOT Logic:\tyes = ",yes,"\t!yes = ", result)
	result = (false && false)
	fmt.Println()
	fmt.Println()
	fmt.Println("AND logic false && false =", result) 
}
/*
AND Logic:	yes && no	 false
OR Logic:	yes || no	 true
NOT Logic:	yes =  true 	!yes =  false


AND logic false && false = false
*/
