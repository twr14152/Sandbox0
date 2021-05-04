//multiple return values

package main

import (
	"fmt"
)

func cube(num int) (string, int, int){
	return "Result", num,(num * num * num)
}

func main() {
	_,b,c := cube(7)
	fmt.Println(b, "Cubed =", c)

	
}

/*
7 Cubed = 343
*/
