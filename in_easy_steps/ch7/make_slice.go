package main

import (
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 3)
	slice2[0] = 11
	slice2[1] = 12
	slice2[2] = 13
	fmt.Print("This is slice1: ")
	describe(slice1)
	fmt.Println()
	fmt.Print("This is slice2: ")
	describe(slice2)
	fmt.Println()
	fmt.Println("This is slice1 with slice2[0] appended which is 11")
	//append(target_slice, item_to_add)
	slice1 = append(slice1, slice2[0])
	describe(slice1)
	fmt.Println("Copy slice1[1:] into slice1[0]")
	//copy(dst[], src[])
	copy(slice1[0:], slice1[1:])
	slice1 = slice1[:len(slice1)-1]
	//The capacity will not go back to original size when elements are removed
	describe(slice1)
	fmt.Println()
	fmt.Println("Copy slice2[2:] or 13 into slice1[4:] in the 4th spot")
	copy(slice1[4:], slice2[2:])
	describe(slice1)

}

func describe(slice1 []int) {
	fmt.Printf("\n%v Length: %v ", slice1, len(slice1))
	fmt.Printf("Capacity: %v \n", cap(slice1))
}
/*
This is slice1: 
[1 2 3 4 5 6 7 8 9 10] Length: 10 Capacity: 10 

This is slice2: 
[11 12 13] Length: 3 Capacity: 3 

This is slice1 with slice2[0] appended which is 11

[1 2 3 4 5 6 7 8 9 10 11] Length: 11 Capacity: 20 
Copy slice1[1:] into slice1[0]

[2 3 4 5 6 7 8 9 10 11] Length: 10 Capacity: 20 

Copy slice2[2:] or 13 into slice1[4:] in the 4th spot

[2 3 4 5 13 7 8 9 10 11] Length: 10 Capacity: 20 
*/
