package main

import "fmt"

func main() {
	nums := make(chan int)
	sqrs := make(chan int)

	go count(nums)
	go square(nums, sqrs)

	for i := 1; i <= 10; i++ {
		num := <-sqrs
		fmt.Printf("%v x %v = %v\n", i, i, num)
	}
}

func count(nums chan<- int) {
	for i := 1; i <= 10; i++ {
		nums <- i
	}
	close(nums)
}

func square(nums <-chan int, sqrs chan<- int) {
	for i := 1; i <= 10; i++ {
		num := <-nums
		sqrs <- num * num
	}
	close(sqrs)
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run synchronize.go 
1 x 1 = 1
2 x 2 = 4
3 x 3 = 9
4 x 4 = 16
5 x 5 = 25
6 x 6 = 36
7 x 7 = 49
8 x 8 = 64
9 x 9 = 81
10 x 10 = 100
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
