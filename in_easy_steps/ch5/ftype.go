package main

import "fmt"

func main() {
	type adder func(int, int) int

	var add adder = func(a int, b int) int {
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
	fmt.Println("Doubled: ", twice(6, 2)*2)
	div := func(a int, b int) int {
		return (a + b) / 2
	}
	return div
}

/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch5 $ go run ftype.go
Added:  8
Doubled:  16
Divided:  4
add type: main.adder
dub type: func(func(int, int) int) func(int, int) int
div type: func(int, int) int
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch5 $
*/
