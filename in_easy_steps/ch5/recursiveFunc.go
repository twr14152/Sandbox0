//Recursive funcs

package main

import (
	"fmt"
)

func countDn(num int) {
	if num < 1 {
		fmt.Println("Blast off!")
	} else {
		fmt.Println("Countdown", num)
		num--
		countDn(num)
	}
}

func main() {
	countDn(10)

}

/*
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
Blast off!
*/
