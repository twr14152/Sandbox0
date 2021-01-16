package main

import (
	"fmt"
)

// wages
var name1_wage float64 = 21
var name2_wage float64 = 20
var name3_wage float64 = 20
var name4_wage float64 = 19

//overtime rate
var name1_ot float64 = (name1_wage / 2) + name1_wage
var name2_ot float64 = (name2_wage / 2) + name2_wage
var name3_ot float64 = (name3_wage / 2) + name3_wage
var name4_ot float64 = (name4_wage / 2) + name4_wage

func wageCalc(hours, wage, ot_wage float64) float64 {
	if hours <= 40 {
		wage_total := hours * wage
		return wage_total
	} else {
		overtime := hours - 40
		ot_total := overtime * ot_wage
		wage_total := 40 * wage
		total_wage := wage_total + ot_total
		return total_wage
	}
}

func main() {
	var name1_hrs float64
	var name2_hrs float64
	var name3_hrs float64
	var name4_hrs float64
	fmt.Println("")
	fmt.Print("Number of hours for name1_hrs: ")
	fmt.Scanf("%f", &name1_hrs)
	fmt.Print("Number of hours for name2_hrs: ")
	fmt.Scanf("%f", &name2_hrs)
	fmt.Print("Number of hours for name3_hrs: ")
	fmt.Scanf("%f", &name3_hrs)
	fmt.Print("Number of hours for name4_hrs: ")
	fmt.Scanf("%f", &name4_hrs)

	name1 := wageCalc(name1_hrs, name1_wage, name1_ot)
	name2 := wageCalc(name2_hrs, name2_wage, name2_ot)
	name3 := wageCalc(name3_hrs, name3_wage, name3_ot)
	name4 := wageCalc(name4_hrs, name4_wage, name4_ot)

	fmt.Println("Name1: ", name1)
	fmt.Println("Name2: ", name2)
	fmt.Println("Name3: ", name3)
	fmt.Println("Name4: ", name4)
	fmt.Println("")

  fmt.Println("Total Amount: ", (name1 +  name2 +  name3 + name4))

}
