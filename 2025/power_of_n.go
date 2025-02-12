package main

import "fmt"

func power(base, exp int) int {
        result:= 1
        for i := 0; i < exp; i++ {
                result *= base
        }
        return result
}

func main() {
        a := 0;
        b := 0;
        fmt.Print("Enter base number: ")
        fmt.Scanf("%d", &a)
        fmt.Print("Enter your exponent: ")
        fmt.Scanf("%d", &b)
        ans := power(a, b)
        fmt.Printf("The answer to %d^%d = %d\n", a, b, ans)
}
/*
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder$ go run power_of_n.go
Enter base number: 7
Enter your exponent: 4
The answer to 7^4 = 2401
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder$
*/
