package main

import (
	"fmt"
	"time"
)

func count(msg string, c chan string) {
	for i := 0; i < 3; i++ {
		c <- msg + " at " + time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go count("Message", c)
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run channel.go 
Message at 26:48
Message at 26:49
Message at 26:50
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
