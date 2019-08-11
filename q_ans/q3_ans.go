package main

import (
	"fmt"
)

var num uint64 = 900
var i uint64

func find_bpn(num uint64) {
	i = 2
	for i < num {
		if num%i == 0 {
			num = num / i
		} else {
			i += 1
		}
	}
	fmt.Println("BPN: ", i)
}

func main() {
	find_bpn(num)
}
