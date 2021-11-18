package main

import (
	"fmt"
)

func main() {

	//fmt.Println("Hey! how you doing")

	a := 10
	b := 15

	fmt.Printf(" %d,%d ", a, b)
	a, b = b, a
	fmt.Printf(" %d,%d ", a, b)

}
