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
Copy directory "verify/verify.go" into go/src
go install verify
*/

