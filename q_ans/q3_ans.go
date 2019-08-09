package main

import (
	"fmt"
)

var num uint64 = 900
var i uint64

func bpf(num uint64) {
	i = 2
	for i < num {
		if num%i == 0 {
			num = num / i
			i = 2
		} else {
			i += 1
		}
	}
	fmt.Println("BPF: ", i)
}

func main() {
	bpf(num)
}
