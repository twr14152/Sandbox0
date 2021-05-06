package main

import (
	"fmt"
)

type employee struct {
	id int
	name string
	dept string
}


func main() {
	var coder employee
	coder.id = 001
	coder.name = "Jack"
	coder.dept = "I.T"
	
	clerk := employee{name:"Ben", dept:"Payroll", id:002}
	
	fmt.Println(coder)
	fmt.Println(clerk)
	fmt.Printf("\n%v works in %v\n", coder.name, coder.dept)
	fmt.Printf("\n%v works in %v\n", clerk.name, clerk.dept)
}
/*
{1 Jack I.T}
{2 Ben Payroll}

Jack works in I.T

Ben works in Payroll

*/
