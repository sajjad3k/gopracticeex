package main

import "math/rand"

type Employee interface {
	age() int
	language() string
}

func (e Engineer) language() string {

	return e.name + "it is"
}

func (e Engineer) age() int {
	return rand.Intn(44)
}

type Engineer struct {
	name string
}

func main() {

	var prog []Employee

	sing := Engineer{name: "singh"}

	prog = append(prog, sing)

}
