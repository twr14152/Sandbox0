package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-results
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Print("Worker ", id, " ran Job ", job)
		fmt.Printf("...%v x %v = %v\n", job, job, job*job)
		time.Sleep(1 * time.Second)
		results <- job
	}
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ go run workers.go 
Worker 3 ran Job 1...1 x 1 = 1
Worker 2 ran Job 3...3 x 3 = 9
Worker 1 ran Job 2...2 x 2 = 4
Worker 3 ran Job 4...4 x 4 = 16
Worker 2 ran Job 5...5 x 5 = 25
Worker 1 ran Job 6...6 x 6 = 36
Worker 3 ran Job 7...7 x 7 = 49
Worker 2 ran Job 8...8 x 8 = 64
Worker 1 ran Job 9...9 x 9 = 81
Worker 3 ran Job 10...10 x 10 = 100
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch11 $ 
*/
