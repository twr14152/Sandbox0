package main

import (
	"fmt"
)

type dog interface {
	speak() string
	move() string
}

type human interface {
	speak() string
	move() string
}

type creature interface {
	dog
	human
}

type bulldog struct{}

type lab struct{}

type person struct{}

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

func (person) speak() string {
	return "Damn it... The dogs got out!!"
}

func (person) move() string {
	return "Chasing the dogs!"
}
func doStuff(c creature) {
	fmt.Printf("\n%v\n", c.speak())
	fmt.Printf("%v\n\n", c.move())
}
func main() {
	var Jack bulldog
	var Gibson lab
	var Todd person
	doStuff(Jack)
	doStuff(Gibson)
	doStuff(Todd)
}

/*
Bark Bark I'm a bulldog..Bark!
Dog spins in circles..Then takes a dump


I just bark until i get what I want.
runs with butt unusually low to the ground


Damn it... The dogs got out!!
Chasing the dogs!

*/
