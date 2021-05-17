//https://play.golang.org/p/svQ4tcdXB59

package main

import (
	"fmt"
	"time"
)

func main() {
  //Example in the book was missing a paramater added 0,
	dt := time.Date(2025, time.January, 1, 12, 0, 0, 0, time.Local)
	fmt.Printf("\nDateTime: %v\n\n", dt)
	dt = dt.AddDate(2, 6, 3)
	fmt.Printf("DateTime: %v\n\n", dt)
	layout := "2006-Jan-02 03:04PM"
	str := "2030-Dec-25 12:30AM"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Parsed DateTime: %v\n", t)
	}

}

/*
DateTime: 2025-01-01 12:00:00 +0000 UTC

DateTime: 2027-07-04 12:00:00 +0000 UTC

Parsed DateTime: 2030-12-25 00:30:00 +0000 UTC
*/
