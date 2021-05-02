package main

import (
	"fmt"
)

func main() {
	var zero, num, max int = 0,0,1
	var up, dn byte = 'A', 'a'
	result := (num == zero)
	fmt.Println("Equality:\tnum == zero\t", result)
	result = (up == dn)
	fmt.Println("Equality:\tup == dn\t", result)
	result = (max != zero)
	fmt.Println("Inequality:\tmax != zero\t", result)
	result = (zero > max)
	fmt.Println("Greater:\tzero > max\t", result)
	result = (max <= zero)
	fmt.Println("Less or Equal:\tmax <= zero\t", result)	
}
/*
Equality:	num == zero	 true
Equality:	up == dn	 false
Inequality:	max != zero	 true
Greater:	zero > max	 false
Less or Equal:	max <= zero	 false
*/
