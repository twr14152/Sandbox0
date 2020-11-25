package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    counts[input.Text()]++
  }
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
/*
 cat test.txt 
This is a test
This is a test
------
 ./main < test.txt 
2   This is a test
*/
