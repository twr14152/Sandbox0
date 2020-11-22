package main

import (
	"fmt"
	"time"
)

func count(item string) {
	start := time.Now()
	for i := 1; i <= 3; i++ {
		fmt.Printf("%v %v", i, item)
		time.Sleep(1 * time.Second)
	}
	stop := time.Now()
	fmt.Printf("\nTime it took to run function was %v.\n", stop.Sub(start))
	fmt.Println()
}

func main() {
	fmt.Println("\nLine-by-Line Execution...")
	count("Rabbit ")
	count("Turtle ")
	fmt.Println("\nConcurrent Execution...")
	go count("Rabbit ")
	count("Turtle ")
}

/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run goroutine.go 

Line-by-Line Execution...
1 Rabbit 2 Rabbit 3 Rabbit 
Time it took to run function was 3.00081512s.

1 Turtle 2 Turtle 3 Turtle 
Time it took to run function was 3.000523582s.


Concurrent Execution...
1 Turtle 1 Rabbit 2 Turtle 2 Rabbit 3 Rabbit 3 Turtle 
Time it took to run function was 3.000578413s.


Time it took to run function was 3.00111015s.

pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
