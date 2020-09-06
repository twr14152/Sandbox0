package main

import (
	"fmt"
)

var (
	base_rate     float64
	regularHrs    = 40.0
	overtimeHrs   float64
	overtime_rate float64
	regularWages  float64
	overtimeWages float64
	total_wages   float64
)

func main() {
	fmt.Println("Whats your base hourly rate for 40hrs?")
	fmt.Scanln(&base_rate)
	fmt.Println("You entered your base pay as $", base_rate, "per hour")
	fmt.Println("How many hours of overtime did you work this week?")
	fmt.Scanln(&overtimeHrs)
	fmt.Println("You entered ", overtimeHrs, "hrs of over time this week.")

	overtime_rate = base_rate * 1.5
	regularWages = base_rate * regularHrs
	overtimeWages = overtime_rate * overtimeHrs
	total_wages = regularWages + overtimeWages

	fmt.Println("Your weekly wage for 40hrs is $", regularWages)
	fmt.Println("Your overtime amount this week is $", overtimeWages)
	fmt.Println("You earned $", total_wages, " this week.")

}
/*
$ go run wageCalc.go 
Whats your base hourly rate for 40hrs?
21.0
You entered your base pay as $ 21 per hour
How many hours of overtime did you work this week?
8.5
You entered  8.5 hrs of over time this week.
Your weekly wage for 40hrs is $ 840
Your overtime amount this week is $ 267.75
You earned $ 1107.75  this week.
$
*/
