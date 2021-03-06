//Pass by Reference

package main

import (
	"fmt"
)

func square(num *int) {
	fmt.Println("\t\tReceived Copy:", num)
	*num *=*num
	fmt.Println("\t\tModified Copy:", num)
}

func main() {
	num := 5
	fmt.Println("Originl:", num)
	square(&num)
	fmt.Println("Original:", num)
	
}

/*
Originl: 5
		Received Copy: 0xc000018050
		Modified Copy: 0xc000018050
Original: 25
*/
