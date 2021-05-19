package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
  var num int = rand.Intn(20)+1
  var guess int = 0
  var flag bool = true
  fmt.Print("Guess my number its between 1-20:")
  for flag {
    _, err := fmt.Scan(&guess)
      if err != nil {
        fmt.Println(err)
      } else if guess > num {
        fmt.Println("Too High, try again:")
      }else if guess < num {
        fmt.Println("Too Low, try again:")
      }else if guess == num {
        fmt.Println("Correct - My number is", num)
        flag = false
      }

  }
}
/*
 ./main
Guess my number its between 1-20:10
Too High, try again:
9
Correct - My number is 9
 


 ./main
Guess my number its between 1-20:10
Too Low, try again:
12
Too Low, try again:
13
Too Low, try again:
15
Correct - My number is 15
 
*/
