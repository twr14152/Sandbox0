package main

import (
	"fmt"
    "io"
	"net/http"
    "strings"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
    if strings.HasPrefix(url, "http://") {
      fmt.Println(url)
      resp, err := http.Get(url)
		  if err != nil {
			  fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		  	os.Exit(1)
	  	}
		  b, err := io.Copy(os.Stdout, resp.Body)
		  resp.Body.Close()
	  	if err != nil {
		  	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
    } else {
      url = "http://" + url
     fmt.Println("This is whats getting tested: ", url)
      resp, err := http.Get(url)
		  if err != nil {
		  	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		  	os.Exit(1)
	  	}
	  	b, err := io.Copy(os.Stdout, resp.Body)
	  	resp.Body.Close()
	  	if err != nil {
	  		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
}
