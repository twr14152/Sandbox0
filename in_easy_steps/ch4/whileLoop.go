package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 5 {
		fmt.Println("Loop iteration", counter)
		counter++
	}
	i := 10
	// Infinte loop
	for {
		fmt.Println("\t\t\tCountdown", i)
		i--
		if i < 1 {
			fmt.Println("\t\t\tLift Off!")
			break
		}
	}
}
/*
Loop iteration 1
Loop iteration 2
Loop iteration 3
Loop iteration 4
Loop iteration 5
			Countdown 10
			Countdown 9
			Countdown 8
			Countdown 7
			Countdown 6
			Countdown 5
			Countdown 4
			Countdown 3
			Countdown 2
			Countdown 1
			Lift Off!

*/
