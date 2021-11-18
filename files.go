package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	s, err := ioutil.ReadFile("hello.go")

	if err != nil {
		fmt.Print(err)
	}

	val := string(s)

	fmt.Println(val)

}
