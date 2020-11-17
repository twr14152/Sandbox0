package main

import (
	"fmt"
)

type member struct {
	firstName string
	lastName  string
}

type article struct {
	title string
	body  string
	member
}

func (m member) fullName() string {
	return m.firstName + " " + m.lastName
}

func (a article) content() {
	fmt.Println("Title:", a.title)
	fmt.Println("Content:", a.body)
	fmt.Printf("Author: %v\n\n", a.fullName())
}

func main() {
	member1 := member{
		"Jack",
		"Riemenschneider",
	}
	article1 := article{
		"Object Oriented Programming",
		"In Go, Composition emulate Inheritance",
		member1,
	}
	article1.content()
	article2 := article{
		"Next title blah blah",
		"Some things else blah blah coming next",
		member1,
	}
	article2.content()
	fmt.Println()
}





/*
Title: Object Oriented Programming
Content: In Go, Composition emulate Inheritance
Author: Jack Riemenschneider

Title: Next title blah blah
Content: Some things else blah blah coming next
Author: Jack Riemenschneider
*/
