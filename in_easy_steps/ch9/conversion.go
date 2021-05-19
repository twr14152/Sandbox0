package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "I can resist eveything except temptation."
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Title(strings.ToLower(str)))

	str = "47 "
	fmt.Printf("\n%v Type: %T, Length: %v\n", str, str, len(str))
	str = strings.Trim(str, " ")
	fmt.Printf("\n%v Type: %T, Length: %v\n", str, str, len(str))

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v Type: %T\n", num, num)
	}
}

/*
I CAN RESIST EVEYTHING EXCEPT TEMPTATION.
i can resist eveything except temptation.
I Can Resist Eveything Except Temptation.

47  Type: string, Length: 3

47 Type: string, Length: 2
47 Type: int
*/
