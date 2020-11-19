package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("Argument %v: %v \t", i, v)
		fmt.Printf("index: %v\t",i)
		fmt.Printf("value: %v\n",v)
	}

	fmt.Println("\nLast Argument:", os.Args[len(os.Args)-1])
}

/*
 go run main.go test1 test2 test3 test4 test5
Argument 0: /tmp/go-build811755167/b001/exe/main    index: 0   value: /tmp/go-build811755167/b001/exe/main
Argument 1: test1   index: 1    value: test1
Argument 2: test2   index: 2    value: test2
Argument 3: test3   index: 3    value: test3
Argument 4: test4   index: 4    value: test4
Argument 5: test5   index: 5    value: test5

Last Argument: test5
 
*/
