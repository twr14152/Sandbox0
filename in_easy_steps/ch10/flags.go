package main

import (
  "fmt"
  "flag"
)

func main() {
  txt := flag.String("txt", "C#", "A string")
  num := flag.Int("num", 8, "An integer")
  sta := flag.Bool("sta", false, "A boolean")

  flag.Parse()
  fmt.Println()
  fmt.Println("Text:", *txt)
  fmt.Println("Number:", *num, " Status:", *sta)

}
/*
go run flags.go

Text: C#
Number: 8  Status: false

go run flags.go -txt=test_flag -num=2020 -sta=true

Text: test_flag
Number: 2020  Status: true
*/
