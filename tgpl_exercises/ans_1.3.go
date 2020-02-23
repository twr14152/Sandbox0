package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timer1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	totalTime1 := time.Since(timer1)
	fmt.Println("Time to run ", totalTime1)

	timer2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	totalTime2 := time.Since(timer2)
	fmt.Println("Time to run ", totalTime2)
}
