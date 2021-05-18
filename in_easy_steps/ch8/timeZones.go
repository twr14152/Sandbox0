package main

import (
	"fmt"
	"time"
)

func main() {
	zone := "All"
	loc, _ := time.LoadLocation("America/New_York")
	dt := time.Now().In(loc)

	abbr, offset := dt.Zone()
	if offset < 1 {
		offset = offset * -1
	}
	offset /= 60
	switch {
	case abbr == "EST" && offset == 300:
		zone = "East Coast"
	case abbr == "EDT" && offset == 240:
		zone = "East Coast"
	}
	fmt.Printf("\nTZ: %v Offset Minutes: %v\n", abbr, offset)
	fmt.Println("Welcome to", zone, "Visitors!")

}
/*
TZ: EST Offset Minutes: 300
Welcome to East Coast Visitors!
*/
