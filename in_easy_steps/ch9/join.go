package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2 := "String one", " and string two."
	str := s1 + s2

	chars, err := fmt.Printf("\n%v\n", str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bytes Written: ", chars)
		fmt.Println("String Length: ", len(str))
	}
	arr := []string{"\nNot", "sure", "where this is going?"}
	ast := strings.Join(arr, " ")
	fmt.Println(ast)
	if ast[0] == 10 && ast[1] == 9 {
		fmt.Printf("1st Char: ASCII %v Newline\n", ast[0])
		fmt.Printf("2nd Char: ASCII %v Tab\n", ast[1])
	}
	fmt.Printf("This line is %v chars long", len("This line is %v chars long"))
}


/*
String one and string two.
Bytes Written:  28
String Length:  26

Not sure where this is going?
This line is 26 chars long
*/
