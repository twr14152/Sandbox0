package main

import (
	"fmt"
)
	
func main() {
	sum := 2 * 3 + 4 - 5
	fmt.Printf("Default Order: %v\n", sum)
	sum = 2*((3+4)-5)
	fmt.Printf("Forced Order: %v\n\n", sum)
	sum = 7%3*2
	fmt.Printf("Default Order: %v\n", sum)
	sum = 7%(3*2)
	fmt.Printf("Forced Order: %v\n\n", sum)
}
/*
Precedence Levels 
======================
**Can supercede these by using ()
**********************
Top-Down L-R
5	* / % << >> &
4	+ - | ^
3	== != < <= >=
2	&&
1	||
**********************

Default Order: 5
Forced Order: 4

Default Order: 2
Forced Order: 1

*/
