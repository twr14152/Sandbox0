/* Example 1 
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
*/
/* Example 2
package main

import (
  "fmt"
  "bufio"
  "os"
)
/* Example 2
func main() {
  counts := make(map[string]int)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
      }
    }
    for line, n := range counts {
      if n > 1 {
        fmt.Printf("%d\t%s\n", n, line)
      }
    }
}
func countLines(f *os.File, counts map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    counts[input.Text()]++
  }
}

/*
 ./main < test.txt
2   This is a test
3   test
 
*/

/* Example 3 */
package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)
func main() {
  counts := make(map[string]int)
  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)
    //data is a byte slice
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
      continue
    }
    // convert byte slice to a string
    for _, line := range strings.Split(string(data), "\n") {
      counts[line]++
    }
    for line, n := range counts {
      if n > 1 {
        fmt.Printf("%d\t%s\n", n, line)
      }
    }
  }
}
/*
 ./main test.txt
2   This is a test
3   test
 
*/
