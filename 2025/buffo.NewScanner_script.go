package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
)

//Use bufio.NewScanner for line by line input
//Use bufio.NewReader for more granular line reading and manipulation

func main() {
        fmt.Println("Enter a sentence and \"quit\" to end...")
        fmt.Printf(":>>")
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
                line := scanner.Text()
                fmt.Println(line)
                if line == "quit" {
                        fmt.Println("Good bye...")
                        break
                }
                fmt.Print("Enter something again>>")
        }
        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }
}
