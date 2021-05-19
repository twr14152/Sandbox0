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
  //ioutil has been deprecated
	txt, err := ioutil.ReadFile("ioutilfile.txt")
	check(err)
	fmt.Println(string(txt))

	file, err := os.Open("osOpenfile.txt")
	check(err)
	defer file.Close()
  b := make([]byte,142)
  fmt.Println(file.Read(b))
  fmt.Println(string(b))
  pos, err := file.Seek(15, 0)
	check(err)

	slice := make([]byte, 15)
	nb, err := file.Read(slice)
	check(err)

	fmt.Printf("\n%v bytes @ %v:", nb, pos)
	fmt.Printf("%v\n", string(slice[:nb]))
}

/*
 go build
 ./main
This file is being openned by ioutil. This functionality has been deprecated by go maintainers.
141 <nil>
This another means of reading file. 
os.Open()
Its considered a more ganular means.
This is a test to see how many characters are reported...

15 bytes @ 15:ans of reading 
 
*/
