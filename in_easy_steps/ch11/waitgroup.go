package main

import (
	"fmt"
	"sync"
	"time"
)

func report(i int, wg *sync.WaitGroup) {
	fmt.Printf("\nGoroutine %v Started", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("\n\t\t\tGoroutine %v Ended", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go report(i, &wg)
	}
	wg.Wait()
	fmt.Println()
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run waitgroup.go 

Goroutine 3 Started
Goroutine 2 Started
Goroutine 1 Started
			Goroutine 3 Ended
			Goroutine 2 Ended
			Goroutine 1 Ended
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
