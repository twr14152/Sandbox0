// Create temp file using the soon to be deprecated (Go1.16)ioutil package
package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func check (err error) {
  if err != nil {
    fmt.Println(err)
  }
}
func main() {
  tmpFile, err := ioutil.TempFile("", "Data-*")
  check(err)
  fmt.Printf("Created file:\n%v\n", tmpFile.Name())
  nb, err := tmpFile.WriteString("Go is the shit!!\n")
  check(err)
  txt, err := ioutil.ReadFile(tmpFile.Name())
  check(err)
  fmt.Printf("\nRead: %v bytes - %v\n", nb, string(txt))
  tmpFile.Close()

  os.Remove(tmpFile.Name())
  _, err = os.Stat(tmpFile.Name())
  check(err)
}
/*
 ./main
Created file:
/tmp/Data-064395144

Read: 17 bytes - Go is the shit!!

stat /tmp/Data-064395144: no such file or directory
 

*/
