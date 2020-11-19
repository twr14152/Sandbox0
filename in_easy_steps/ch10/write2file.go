//So far it looks like io/ioutil is good for creating and reading files
//os.O_APPEND|os.O_WRONLY used for appending files using filname.Write options


package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	bgp_conf := []byte("\nrouter bgp 65535 \n  network 10.0.0.0 mask 255.255.255.0\n  neighbor 10.0.0.3 remote-as 65535\n")
	err := ioutil.WriteFile("testfile001.txt", bgp_conf, 0644)
	check(err)
	file, err := os.OpenFile("testfile001.txt", os.O_APPEND|os.O_WRONLY, 0644)
	//data, err := ioutil.ReadFile("testfile001.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))
	check(err)

	vty_conf := []byte("line vty 0 4\n  transport input ssh \n  authentication local \n ")
	app_txt, err := file.Write(vty_conf)
	check(err)
	data, _ := ioutil.ReadFile("testfile001.txt")
	fmt.Println("Final config:", string(data))
	fmt.Printf("\nAppended: %v bytes - %v", app_txt, string(vty_conf[:app_txt]))
	defer file.Close()
}
/*
pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ go run write2file.go 
Final config: 
router bgp 65535 
  network 10.0.0.0 mask 255.255.255.0
  neighbor 10.0.0.3 remote-as 65535
line vty 0 4
  transport input ssh 
  authentication local 
 

Appended: 61 bytes - line vty 0 4
  transport input ssh 
  authentication local 
 pi@RaspPi4:~/Coding/Go_folder/in_easy_steps/ch10 $ 
*/
