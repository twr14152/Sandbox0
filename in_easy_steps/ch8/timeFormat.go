//https://play.golang.org/p/Cv0q1XvSazW
package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println("\nDefault Format:", dt)
	fmt.Println("Unix Format:", dt.Format(time.UnixDate))
	fmt.Println("ANSIC Format:", dt.Format(time.ANSIC))
	fmt.Println("RFC3339 Format:", dt.Format(time.RFC3339))
	fmt.Println("Custom Format:", dt.Format("January 2, 2006[Monday]"))

	fmt.Println("\nUS Format:", dt.Format("October 8, 2006"))
	fmt.Println("UK Format:", dt.Format("8 October, 2006"))

	fmt.Println("\nTime 12-Hour:", dt.Format(time.Kitchen))
	fmt.Println("Time 24-Hour:", dt.Format("15:04"))

}

/*
Default Format: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
Unix Format: Tue Nov 10 23:00:00 UTC 2009
ANSIC Format: Tue Nov 10 23:00:00 2009
RFC3339 Format: 2009-11-10T23:00:00Z
Custom Format: November 10, 2009[Tuesday]

US Format: October 8, 2009
UK Format: 8 October, 2009

Time 12-Hour: 11:00PM
Time 24-Hour: 23:00
*/
