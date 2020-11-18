
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(59)
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
	str := "\nYour Six Lucky Number:"
	for i := 0; i < 6; i++ {
		str += strconv.Itoa(nums[i])
		if i != 5 {
			str += "-"
		}
	}
	fmt.Println(str)

}

/*
Your Six Lucky Number:2-6-51-13-18-33
*/
