//bufio is good for inputing larger streams of data

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("\nEnter Text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		fmt.Println(scanner.Text())
	}
	fmt.Print("\nEnter Text:")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
  fmt.Printf("Words Type: %T\t Value: %v\n", words, words)
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		for i, v := range words {
			fmt.Printf("%v: %v\n", i, v)
		}

	}
}
/*
 ./main

Enter Text:This is a test
This is a test

Enter Text:This is a test
Words Type: []string     Value: [This is a test]
0: This
1: is
2: a
3: test
 
