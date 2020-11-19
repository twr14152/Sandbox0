package main

import (
	"fmt"
	"os"
)

func main() {

	for i, v := range os.Args {
		fmt.Printf("Argument %v: %v \n", i, v)
	}

	fmt.Println("\nLast Argument:", os.Args[len(os.Args)-1])
}

/*
 go run main arg1 arg2 arg3 arg4 arg5
Argument 0: /tmp/go-build981861993/b001/exe/main 
Argument 1: arg1 
Argument 2: arg2 
Argument 3: arg3 
Argument 4: arg4 
Argument 5: arg5 

Last Argument: arg5
 
*/
