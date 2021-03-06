package main

import "fmt"

type car struct {
	color string
	body  string
}

/*
When you call a method, Go automatically handles the conversion between values and pointers,
so the receiver here could be ( c *car ) if you wanted the method to be able to modify the 
original struct fields.
*/

func (c car) accelerate() string {
	return "accelerating-->"
}

func main() {
	porsche := car{color: "blue", body: "coupe"}
	bentley := car{color: "green", body: "soloon"}
	fmt.Println("Porsche paint is", porsche.color)
	fmt.Println("Porsche style is", porsche.body)
	fmt.Println("Porsche is", porsche.accelerate())
	fmt.Println("Bentley paint is", bentley.color)
	fmt.Println("Bentley style is", bentley.body)
	fmt.Println("Bentley is", bentley.accelerate())
}

/*
Porsche paint is blue
Porsche style is coupe
Porsche is accelerating-->
Bentley paint is green
Bentley style is soloon
Bentley is accelerating-->
*/
