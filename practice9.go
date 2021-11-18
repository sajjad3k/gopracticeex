package main

import "fmt"

//type MyInt int

func IsArmstrong(n int) int {

	arm := 0

	for n > 0 {

		tmp := n % 10
		arm = arm + (tmp * tmp * tmp)
		n = n / 10
	}

	return arm
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 int = 371
	fmt.Println(IsArmstrong(num1))
}
