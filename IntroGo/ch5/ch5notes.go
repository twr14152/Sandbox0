package main

import "fmt"

func main() {
	//This is an array
	var x [5]int
	x[4] = 100
	fmt.Println("This is an array ", x)
	fmt.Println("Now we are going insert values in to array and sum them")
	x[0] = 98
	x[1] = 1000
	x[2] = 234
	x[3] = 5406
	x[4] = 6387

	var total int = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
		fmt.Printf("The value of %v\t, %v\n", i, total)
	}
	fmt.Println("Total =", total)
	fmt.Println("Avg value of x is", total / len(x))
	total = 0
	// This is a slice
	y := []float64{0,1,2,3,4,5,6,7,8,9,10}
	for _, value := range y {
		//This is type conversion value is float64 total is int
		total += int(value)
	}
	fmt.Println("Total value of y =  ",  total)
	fmt.Println("Avg value of y is = ", total/(len(y)))
	// make([]type, len, cap)
	z := make([]float64, 8, 10)
	fmt.Println(z)
	z[0] = 10.0
	fmt.Println("Assign 10.0 to index 0", z)
	// copy slice y to aa
	aa := make([]float64, len(y))
	// to copy you need the slices to be the same len. nil will not work
	fmt.Println("slice z = ", y)
	fmt.Println("slice aa = ", aa)
	fmt.Println(copy(aa, y))
	fmt.Println(aa, y)
	//maps require make() otherwise runtime err will occur
	bb := make(map[string]int)
	bb["Jack"] = 5
	bb["Gibson"] = 2
	bb["Luna"] = 1
	fmt.Println(bb)
	fmt.Println(bb["Jack"])
	// nested maps
	Pets := map[string]map[string]string{
		"dogs": map[string]string{
			"bull dog": "Jack",
			"black lab":"Gibson",
		},
		"cats": map[string]string{
			"black kitten": "Luna",
		},
	}
	fmt.Println(Pets)
	fmt.Println(Pets["dogs"])
	fmt.Println(Pets["cats"])
	fmt.Println("name of dogs, bull dog:", Pets["dogs"]["bull dog"])
	fmt.Println("name of cats, black kitten: " , Pets["cats"]["black kitten"])
	//loop through map
	for k, v := range Pets {
		fmt.Println("Loop through keys in maps:\n" ,k)
		fmt.Println("Loop through values in maps:\n", v)
	}
}
