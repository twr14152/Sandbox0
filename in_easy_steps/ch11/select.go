package main

import (
	"fmt"
	"time"
)

func main() {
	c0 := time.After(7 * time.Second)
	c1 := make(chan string)
	c2 := make(chan string)

	go fastCount(c1)
	go slowCount(c2)

Listener:
	for {
		select {
		case done := <-c0:
			fmt.Println("Timed Out at", done.Format("04:05"))
			break Listener
		case msg1 := <-c1:
			fmt.Println("1-Second Message at ", msg1)
		case msg2 := <-c2:
			fmt.Println("\t\t\t\t2-Second Message at ", msg2)
		}
	}
}

func fastCount(c1 chan string) {
	for {
		time.Sleep(1 * time.Second)
		c1 <- time.Now().Format("04:05")
	}
}

func slowCount(c2 chan string) {
	for {
		time.Sleep(2 * time.Second)
		c2 <- time.Now().Format("04:05")
	}
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run select.go 
1-Second Message at  58:23
				2-Second Message at  58:24
1-Second Message at  58:24
1-Second Message at  58:25
				2-Second Message at  58:26
1-Second Message at  58:26
1-Second Message at  58:27
				2-Second Message at  58:28
1-Second Message at  58:28
Timed Out at 58:29
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
