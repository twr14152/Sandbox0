package main

import (
	"fmt"
)

func main() {
	array := [3]string{"BMW", "FORD", "CHEVY"}
	slice := array[:]
	list(slice...)
	slice = append(slice, "Porsche", "Ferrari", "Honda", "Toyota")
	list(slice...)
}

func list(autos ...string) {
	for i, v := range autos {
		fmt.Printf("\n%v. %v", i, v)
	}
	fmt.Println()
}

/*
0. BMW
1. FORD
2. CHEVY

0. BMW
1. FORD
2. CHEVY
3. Porsche
4. Ferrari
5. Honda
6. Toyota
*/
