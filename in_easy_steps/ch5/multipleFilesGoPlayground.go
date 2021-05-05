package main

import (
	"fmt"
	"play.ground/verify"
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
-- go.mod --
module play.ground
-- verify/verify.go --
package verify

import "fmt"

func IsPosInt(num int) (int, error) {
	if num < 1 {
		err := fmt.Errorf("%v not a positive integer", num)
		return -1, err
	}
	return num, nil
}
/*
2 Passed!
1 Passed!
Failed:  0 not a positive integer
Failed:  -1 not a positive integer
Failed:  -2 not a positive integer
*/
