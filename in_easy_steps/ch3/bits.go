package main

import (
	"fmt"
)
	
func main() {
	var flags byte = 10 //binary 1010
	fmt.Printf("Flag #1:%v \n", (flags & 1) > 0)
	fmt.Printf("Flag #2:%v \n", (flags & 2) > 0)
	fmt.Printf("Flag #3:%v \n", (flags & 4) > 0)
	fmt.Printf("Flag #4:%v \n", (flags & 8) > 0)
	fmt.Printf("Flags Value: %08b \t%v \n", flags, flags)
	flags = flags >> 1 // shift right 1
	fmt.Printf("New Value: %08b \t%v \n", flags, flags)
	
}
/*
Flag #1:false 
Flag #2:true 
Flag #3:false 
Flag #4:true 
Flags Value: 00001010 	10 
New Value: 00000101 	5 
*/
