package main

import (
	"fmt"
	//"strings"
)

func main() {
	s1, s2 := "String one", " and string two."
	str := s1 + s2
	//The line below does 3 things
	//Prints output of str
	//assigns number of bytes written to char
	//assigns an error message
	chars, err := fmt.Printf("\n%v\n", str)
	fmt.Println(chars)
	fmt.Println(err)
	
	//The line below does 3 things
	//Prints output of str
	//assigns number of bytes written to char
	//assigns an error message	
	chars, err = fmt.Println(str)
	fmt.Println(chars)
	fmt.Println(err)	
}
/*

String one and string two.
28
<nil>
String one and string two.
27                          
<nil>
*/
