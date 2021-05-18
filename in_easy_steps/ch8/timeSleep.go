//https://play.golang.org/p/Z1HDT2SGPc6
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("\nStarted at:", start.Format("03:45:05"))
	time.Sleep(5 * time.Second)

	finish := time.Now()
	fmt.Println("Finished at:", finish.Format("03:04:05"))
	fmt.Println("\nStart First?:", start.Before(finish))
	fmt.Println("Finish First?:", finish.Before(start))

	diff := finish.Sub(start)
	fmt.Println("\nTime Elapsed:", diff.Round(time.Second))
}
/*
Started at: 11:00:00
Finished at: 11:00:05

Start First?: true
Finish First?: false

Time Elapsed: 5s
*/
