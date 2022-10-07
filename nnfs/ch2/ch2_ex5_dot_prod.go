package main

import "fmt"

func main() {
	a := []float64{1, 2, 3}
	b := []float64{2, 3, 4}

	dot_prod := a[0]*b[0] + a[1]*b[1] + a[2]*b[2]

	fmt.Printf("dot_product of a and b is %f\n", dot_prod)

}

//pi@RaspPi4:~/Coding/Go_folder/nnfs/ch2 $ go run ch2_ex5_dot_prod.go 
//dot_product of a and b is 20.000000
//pi@RaspPi4:~/Coding/Go_folder/nnfs/ch2 $ 
