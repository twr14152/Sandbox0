package main

import (
	"fmt"
)

func main() {
	var num int = 20
	var ptr *int = &num
	fmt.Printf("num value:%v type:%T \n", num, num)
	fmt.Printf("num address:%v type:%T \n\n", ptr, ptr)
	fmt.Printf("num via pointer:%v type:%T \n", *ptr, *ptr)
	fmt.Printf("ptr address:%v type:%T \n\n", &ptr, &ptr)
	*ptr = 100
	fmt.Printf("new num value:%v type:%T \n", num,num)
	fmt.Println()
	fmt.Printf("num = %v\n", num)
	fmt.Printf("&num = %v\n", &num)
	fmt.Println()
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("*ptr = %v\n", *ptr)
	fmt.Printf("&ptr = %v\n", &ptr)
}

/*
num value:20 type:int 
num address:0xc000018050 type:*int 

num via pointer:20 type:int 
ptr address:0xc00000e028 type:**int 

new num value:100 type:int 

num = 100
&num = 0xc000018050

ptr = 0xc000018050
*ptr = 100
&ptr = 0xc00000e028
*/
