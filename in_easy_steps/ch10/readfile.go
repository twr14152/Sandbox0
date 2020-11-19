package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	txt, err := ioutil.ReadFile("testfile001.txt")
	check(err)
	fmt.Println(string(txt))

	file, err := os.Open("testfile001.txt")
	check(err)
	defer file.Close()

	pos, err := file.Seek(50, 0)
	check(err)

	slice := make([]byte, 15)
	nb, err := file.Read(slice)
	check(err)

	fmt.Printf("\n%v bytes @ %v:", nb, pos)
	fmt.Printf("%v\n", string(slice[:nb]))
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ go run readfile.go 
Test file

The writing in the file is for sample only.

config t
interface Gi0/0
ip address 10.0.0.1/32
no shut

router ospf 1
 network 10.0.0.0/32 area 0




15 bytes @ 50:nly.

config t

pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ 
*/
