// Custom function types

package main

import "fmt"

func main() {
	type adder func(int, int) int

	var add adder = func(a int, b int) int {
		fmt.Print("(func add): ")
		return a + b
	}
	fmt.Println("Added: ", add(6, 2))
	div := dub(add)
	fmt.Println("Divided: ", div(6, 2))
	fmt.Printf("add type: %T\n", add)
	fmt.Printf("dub type: %T\n", dub)
	fmt.Printf("div type: %T\n", div)

}

func dub(twice func(int, int) int) func(int, int) int {
	fmt.Println("(func dub): Doubled: ", twice(6, 2)*2)
	div := func(a int, b int) int {
		fmt.Print("(func div): ")
		return (a + b) / 2
	}
	return div
}
/*
(func add): Added:  8
(func add): (func dub): Doubled:  16
(func div): Divided:  4
add type: main.adder
dub type: func(func(int, int) int) func(int, int) int
div type: func(int, int) int
*/
