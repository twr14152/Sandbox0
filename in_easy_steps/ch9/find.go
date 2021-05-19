//https://play.golang.org/p/a7stweU51gn
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I can resist eveything except temptation."
	fmt.Printf("\nFound 'an':%v\n", strings.Contains(str, "an"))
	fmt.Printf("Found 'an' at:%v\n", strings.Index(str, "an"))
	fmt.Printf("Count of 'e':%v\n", strings.Count(str, "e"))
	fmt.Printf("Prefix 'ion', %v\n", strings.HasPrefix(str, "ion"))
	fmt.Printf("Suffix 'ion', %v\n", strings.HasSuffix(str, "ion"))
	fmt.Println(strings.Replace(str, "temptation", "lots of MONEY!!!", 1))
}
/*
Found 'an':true
Found 'an' at:3
Count of 'e':6
Prefix 'ion', false
Suffix 'ion', false
I can resist eveything except lots of MONEY!!!.
*/
