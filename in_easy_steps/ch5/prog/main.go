package main

import (
	"fmt"
	"verify"
)

func main() {
	for i := 2; i >= -2; i-- {
		res, err := verify.IsPosInt(i)
		if err != nil {
			fmt.Println("Failed: ", err)
		} else {
			fmt.Println(res, "Passed!")
		}
	}
}

/*
pi@RaspPi4:~/go/src/prog $ go build
pi@RaspPi4:~/go/src/prog $ ls -l
total 1916
-rw-r--r-- 1 pi pi     219 Nov 15 20:07 main.go
-rwxr-xr-x 1 pi pi 1951308 Nov 15 20:09 prog
pi@RaspPi4:~/go/src/prog $ ./prog 
2 Passed!
1 Passed!
Failed:  0 not a positive integer
Failed:  -1 not a positive integer
Failed:  -2 not a positive integer
pi@RaspPi4:~/go/src/prog $ 
*/
