package main

import (
  "fmt"
  "flag"
)

func main() {
  //setting default values for flags
  //to change values from default ./main -txt=blah -num=0 sta=true
  txt := flag.String("txt", "C#", "A string")
  num := flag.Int("num", 8, "An integer")
  sta := flag.Bool("sta", false, "A boolean")

  flag.Parse()
  fmt.Println()
  fmt.Println("Text:", *txt)
  fmt.Println("Number:", *num) 
  fmt.Println("Status:", *sta)

  if *num > 8 && *num < 21 {
    for i := 0; i < *num; i++ {
      fmt.Println("Num count: ", i)
    }
  }else{
    fmt.Println("The end")
  }

}
/*
go run flags.go

Text: C#
Number: 8  Status: false

go run flags.go -txt=test_flag -num=2020 -sta=true

Text: test_flag
Number: 2020  Status: true

go run flags.go -txt=test_flag -num=10 -sta=true

Text: test_flag
Number: 10
Status: true
Num count:  0
Num count:  1
Num count:  2
Num count:  3
Num count:  4
Num count:  5
Num count:  6
Num count:  7
Num count:  8
Num count:  9

go run flags.go -txt=test_flag -num=9 -sta=true

Text: test_flag
Number: 9
Status: true
Num count:  0
Num count:  1
Num count:  2
Num count:  3
Num count:  4
Num count:  5
Num count:  6
Num count:  7
Num count:  8

go run flags.go -txt=test_flag -num=8 -sta=true

Text: test_flag
Number: 8
Status: true
The end
*/
