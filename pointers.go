package main

import "fmt"

func main() {
	//s := 23
	//k := &s

	l := "hey how you doing"

	/* fmt.Println(s)
	fmt.Println(k)
	fmt.Println(*k) */

	fmt.Println(l[0:5])
	for i := 0; i < len(l); i++ {
		fmt.Printf("%c\n", l[i])
	}

}
