package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(59)
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
	str := "\nYour Six Lucky Number:"
	for i := 0; i < 6; i++ {
		str += strconv.Itoa(nums[i])
		if i != 5 {
			str += "-"
		}
	}
	fmt.Println(str)

}

/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch9 $ go run lucky6.go 

Your Six Lucky Number:57-13-56-51-18-24
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch9 $ go run lucky6.go 

Your Six Lucky Number:9-5-59-41-46-58
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch9 $ go run lucky6.go 

Your Six Lucky Number:40-2-38-42-37-19
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch9 $ 
*/
