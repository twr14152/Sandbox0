package main

import (
	"fmt"
	"time"
)

func count(msg string, c chan string) {
	fmt.Println("Channel Sender")
	for i := 0; i < 3; i++ {
		fmt.Print("Sending: ")
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
Sending: Message at 45:38
Channel Receiver - Open
Sending: Message at 45:39
Channel Receiver - Open
Sending: Message at 45:40
Channel Receiver - Open
Sender closing channel
Channel receiver - Closed
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 

*/
