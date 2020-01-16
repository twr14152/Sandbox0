package main

import (
	"fmt"
)

func reverse_order(n int) int {
	num := 0
	for n > 0 {
		remainder := n % 10
		num *= 10
		num += remainder
		n /= 10
	}
	return num
}

func main() {
	for a := 900; a <= 999; a++ {
		for b := 900; b <= 999; b++ {
			product := a * b
			if product == reverse_order(product) {
				fmt.Println(product)
			} else {
				continue
			}
		}

	}
}
