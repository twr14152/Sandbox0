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
	delete(pets, "Mitts")
	fmt.Println("The length of pets map is", len(pets))
	fmt.Println("The items in the map are:", pets)
}
/*
Pets:map[Evee:tabby Gibby:Lab Jack:Bulldog Luna:tabby Mitts:Manecoon]
Jack is an american Bulldog
The length of pets map is 5
The length of pets map is 4
The items in the map are: map[Evee:tabby Gibby:Lab Jack:Bulldog Luna:tabby]
*/
