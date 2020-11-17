package main

import (
	"fmt"
)

type dog interface {
	speak() string
	move() string
}

type bulldog struct{}

type lab struct{}

func (bulldog) speak() string {
	return "Bark Bark I'm a bulldog..Bark!"
}

func (bulldog) move() string {
	return "Dog spins in circles..Then takes a dump"
}

func (lab) speak() string {
	return "I just bark until i get what I want."
}

func (lab) move() string {
	return "runs with butt unusually low to the ground"
}

func doStuff(d dog) {
	fmt.Printf("\n%v\n", d.speak())
	fmt.Printf("%v\n\n", d.move())
}
func main() {
	var Jack bulldog
	var Gibson lab
	doStuff(Jack)
	doStuff(Gibson)
}


/*

Bark Bark I'm a bulldog..Bark!
Dog spins in circles..Then takes a dump


I just bark until i get what I want.
runs with butt unusually low to the ground

*/
