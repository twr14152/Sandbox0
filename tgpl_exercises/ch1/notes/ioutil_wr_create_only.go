//Looks like ioutil is good for reading and creating files.
//Use os.OpenFile to append data to existing file.

package main

import (
	"io/ioutil"
	"log"
)

func main() {

	message := []byte("This line will get overwritten")
	message1 := []byte("Does not look like you can append.\nThis is line 2\nline 3\nEtc..\n")
	err := ioutil.WriteFile("hello.txt", message, 0644)
	err = ioutil.WriteFile("hello.txt", message1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}


/*
 ./main
 cat hello.txt 
Does not look like you can append.
This is line 2
line 3
Etc..
 
*/
