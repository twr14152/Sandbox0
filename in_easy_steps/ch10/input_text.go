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
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		for i, v := range words {
			fmt.Printf("%v: %v\n", i, v)
		}

	}
}

/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ go run input_text.go 

Enter Text:This is a test         
This is a test

Enter Text:This is a test again
0: This
1: is
2: a
3: test
4: again
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ 
*/
