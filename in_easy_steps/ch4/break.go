package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("Running i=", i, "j=", j)
			if i == 3 && j == 2 {
				fmt.Println("Continues When i=", i, "and j=", j)
				continue
			}

			if i == 2 && j == 2 {
				fmt.Println("Break When i=", i, "and j=", j)
				break
			}
		}
	}
}

/*
Running i= 1 j= 1
Running i= 1 j= 2
Running i= 1 j= 3
Running i= 2 j= 1
Running i= 2 j= 2
Break When i= 2 and j= 2
Running i= 3 j= 1
Running i= 3 j= 2
Continues When i= 3 and j= 2
Running i= 3 j= 3
*/
