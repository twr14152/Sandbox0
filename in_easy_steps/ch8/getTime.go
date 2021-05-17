//https://play.golang.org/p/TlIkzC7cU3Z
package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Printf("DateTime: %v\n", dt)
	fmt.Printf("DateTime Type: %T\n\n", dt)

	hr := dt.Hour()
	switch {
	case hr < 12:
		fmt.Println("Good Morning")
	case hr < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
	h, mn, s := dt.Clock()
	fmt.Printf("Time: %v:%v,:%v\n", h, mn, s)

	ns := dt.UnixNano()
	ms := ns / 1000000
	fmt.Println("Nanoseconds:", ns)
	fmt.Println("Milliseconds:", ms)

}

/*
DateTime: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
DateTime Type: time.Time

Good Evening
Time: 23:0,:0
Nanoseconds: 1257894000000000000
Milliseconds: 1257894000000
*/
