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
This is not a matching line
test
test
test
------
 ./main < test.txt 
2   This is a test
3   test

----
%d de cimal integer
%x, %o, %b integer in hexade cimal, octal, binar y
%f, %g, %e floating-p oint number: 3.141593 3.141592653589793 3.141593e+00
%t boole an: true or false
%c rune (Unico de co de point)
%s st ring
%q quot ed str ing "abc" or rune 'c'
%v any value in a natural for mat
%T type of any value
%% literal percent sig n (no operand)
*/
