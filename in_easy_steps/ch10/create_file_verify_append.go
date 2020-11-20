package main

import (
  "fmt"
  "os"
  "os/exec"
  "io/ioutil"
)

func execute() {
    out, err := exec.Command("ls", "-l").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Verfying config_file.txt generated")
    output := string(out[:])
    fmt.Println(output)
}

func createConfig() {
  f, _ := os.Create("config_file.txt")
  f.WriteString("config t\n  interface loopback 0\n  ip address 1.1.1.1/32\n")
}
func readConfig() {
    dat, err := ioutil.ReadFile("config_file.txt")
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println(string(dat))
    }
}

func addConfig() {
   f, err := os.OpenFile("config_file.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    newLines := "\nline vty 0 4\n login local\n  transport input ssh\n"
    _, err = fmt.Fprintln(f, newLines)
    if err != nil {
      fmt.Println(err)
    }
}
func main() {
  fmt.Println("Generate config file....\n")
  createConfig()
  execute()
  fmt.Println("Contents of the config:\n\n")
  readConfig()
  addConfig()
  fmt.Println("Contents of the config file after addition:\n\n")
  readConfig()
}

/*
Generate config file....

Verfying config_file.txt generated
total 2360
-rw-r--r-- 1 runner runner      56 Nov 20 13:28 config_file.txt
-rw-r--r-- 1 runner runner      21 Nov 19 01:09 go.mod
-rwxr-xr-x 1 runner runner 2402629 Nov 20 13:28 main
-rw-r--r-- 1 runner runner    1716 Nov 20 13:15 main.go

Contents of the config:


config t
  interface loopback 0
  ip address 1.1.1.1/32

Contents of the config file after addition:


config t
  interface loopback 0
  ip address 1.1.1.1/32

line vty 0 4
 login local
  transport input ssh

*/
