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
	describe(slice1)
	slice1 = append(slice1, slice2[0])
	describe(slice1)
	//copy(dst[], src[])
	copy(slice1[0:], slice1[1:])
	slice1 = slice1[:len(slice1)-1]
	describe(slice1)

}

func describe(slice1 []int) {
	fmt.Printf("\n%v Length: %v ", slice1, len(slice1))
	fmt.Printf("Capacity: %v \n", cap(slice1))
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch7 $ go run make_slice.go 

[1 2 3 4 5 6 7 8 9 10] Length: 10 Capacity: 10 

[1 2 3 4 5 6 7 8 9 10 11] Length: 11 Capacity: 20 

[2 3 4 5 6 7 8 9 10 11] Length: 10 Capacity: 20 
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch7 $ 
*/
