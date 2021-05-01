package main

import "fmt"

func main() {
	num := 100
	pi := 3.1415926536
	fmt.Printf("num: %v type:%T\n", num, num)
	fmt.Printf("pi: %v type:%T\n\n", pi, pi)
	fmt.Printf("%%7d displays %7d\n", num)
	fmt.Printf("%%07d displays %07d\n\n", num)
	fmt.Printf("Pi is approximately %1.10f\n", pi)
	fmt.Printf("Right-aligned %20.3f rounded pi\n", pi)
	fmt.Printf("Left-Aligned %-20.3f rounded pi \n", pi)
}

/*
num: 100 type:int
pi: 3.1415926536 type:float64

%7d displays     100
%07d displays 0000100

Pi is approximately 3.1415926536
Right-aligned                3.142 rounded pi
Left-Aligned 3.142                rounded pi 
*/
