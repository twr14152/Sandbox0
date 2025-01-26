package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"bufio"
)

func main() {
	fmt.Printf("Enter url to search >>")
	url := bufio.NewScanner(os.Stdin)
	url.Scan()
	fmt.Printf("GET http://%v\n",url.Text())
	res, err := http.Get("http://"+url.Text())
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

pi@ra
