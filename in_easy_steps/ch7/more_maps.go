package main

import (
	"fmt"
)

func main() {
	pets := make(map[string]string)
	pets["Jack"] = "Bulldog"
	pets["Gibby"] = "Lab"
	pets["Evee"] = "tabby"
	pets["Luna"] = "tabby"
	pets["Mitts"] = "Manecoon"

	fmt.Printf("Pets:%v\n", pets)
	fmt.Printf("Jack is an american %v\n", pets["Jack"])
	fmt.Printf("The length of pets map is %v\n", len(pets))
	for k, v := range pets {
		fmt.Printf("The animals name is %v and their type is %v\n", k, v)
	}
	delete(pets, "Mitts")
	fmt.Println("The length of pets map is", len(pets))
	fmt.Println("The items in the map are:", pets)
}
/*
Pets:map[Evee:tabby Gibby:Lab Jack:Bulldog Luna:tabby Mitts:Manecoon]
Jack is an american Bulldog
The length of pets map is 5
The animals name is Evee and their type is tabby
The animals name is Luna and their type is tabby
The animals name is Mitts and their type is Manecoon
The animals name is Jack and their type is Bulldog
The animals name is Gibby and their type is Lab
The length of pets map is 4
The items in the map are: map[Evee:tabby Gibby:Lab Jack:Bulldog Luna:tabby]
*/
