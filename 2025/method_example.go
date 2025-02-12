
package main

import (
	"fmt"
)

type FingersLeft int

func (i FingersLeft) LostFingers() int {
	if i < 0 {
		return 0
	}
	return 10 - int(i)
}

func main() {
    var f FingersLeft = 6
	fmt.Printf("You have %v fingers left since you lost %v in the accident.\n", f, f.LostFingers())
}


/*

*/
