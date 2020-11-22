package main

import (
	"fmt"
	"time"
)

func count(msg string, c chan string) {
	fmt.Println("Channel Sender")
	for i := 0; i < 3; i++ {
		fmt.Println("Sending message: ")
		c <- msg + " at " + time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
	close(c)
	fmt.Println("Sender closing channel")
}

func main() {
	c := make(chan string)
	go count("Message", c)
	for {
		fmt.Println("Channel Receiver - Open")
		msg, open := <-c
		if !open {
			fmt.Println("Channel receiver - Closed")
			break
		}
		fmt.Println(msg)
	}
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run channel.go 
Channel Receiver - Open
Channel Sender
Sending message: 
Message at 43:21
Channel Receiver - Open
Sending message: 
Message at 43:22
Channel Receiver - Open
Sending message: 
Message at 43:23
Channel Receiver - Open
Sender closing channel
Channel receiver - Closed
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
