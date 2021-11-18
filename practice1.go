package main

import "fmt"

//type assertions

type dev struct {
	nme string
	ag  int
}

func Getdev(name, age interface{}) dev {

	f := dev{}
	f.nme = name.(string)
	f.ag = age.(int)
	return f
}

func main() {

	var name interface{} = "sajjad"
	var age interface{} = 24

	devy := Getdev(name, age)

	fmt.Println(devy.nme)
	fmt.Println(devy.ag)

}
