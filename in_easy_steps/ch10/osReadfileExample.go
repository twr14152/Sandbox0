// This is the recommended replacement to ioutil thats been deprecated.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("fizzbuzz.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	os.Stdout.Write(data)
}
/*
pi@RaspPi4:~/Coding/Go_folder/misc_stuff $ go run osReadfile_example.go 
[47 47 104 116 116 112 115 58 47 47 112 108 97 121 46 103 111 108 97 110 103 46 111 114 103 47 112 47 55 122 52 122 75 110 100 106 48 75 98 10 10 112 97 99 107 97 103 101 32 109 97 105 110 10 10 105 109 112 111 114 116 32 40 10 9 34 102 109 116 34 10 41 10 10 102 117 110 99 32 109 97 105 110 40 41 32 123 10 9 102 111 114 32 105 32 58 61 32 48 59 32 105 32 60 32 49 48 49 59 32 105 43 43 32 123 10 9 9 105 102 32 105 37 51 32 61 61 32 48 32 38 38 32 105 37 53 32 61 61 32 48 32 123 10 9 9 9 102 109 116 46 80 114 105 110 116 108 110 40 34 102 105 122 122 32 98 117 122 122 32 61 32 34 44 32 105 41 10 9 9 125 32 101 108 115 101 32 105 102 32 105 37 51 32 61 61 32 48 32 123 10 9 9 9 102 109 116 46 80 114 105 110 116 108 110 40 34 70 105 122 122 32 61 32 34 44 32 105 41 10 9 9 125 32 101 108 115 101 32 105 102 32 105 37 53 32 61 61 32 48 32 123 10 9 9 9 102 109 116 46 80 114 105 110 116 108 110 40 34 66 117 122 122 32 61 32 34 44 32 105 41 10 9 9 125 32 101 108 115 101 32 123 10 9 9 9 102 109 116 46 80 114 105 110 116 108 110 40 34 78 111 114 109 97 108 32 61 34 44 32 105 41 10 9 9 125 10 10 9 125 10 125 10]
//https://play.golang.org/p/7z4zKndj0Kb

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz = ", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz = ", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz = ", i)
		} else {
			fmt.Println("Normal =", i)
		}

	}
}
*/
