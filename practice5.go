package main

import (
	"errors"
	"fmt"
)

type stack struct {
	item []Fligh
}

type Fligh struct {
	origin      string
	destination string
	price       int
}

func (s *stack) Isempty() bool {
	if len(s.item) == 0 {
		return true
	}
	return false
}

func (s *stack) push(flight Fligh) {

	s.item = append(s.item, flight)

}

func (s *stack) pop() (Fligh, error) {

	if s.Isempty() == true {
		return Fligh{}, errors.New("stack is empty")
	}

	lastindex := len(s.item) - 1
	flight := s.item[lastindex]
	s.item = s.item[:lastindex]

	return flight, nil
}

func (s *stack) peek() (Fligh, error) {

	if s.Isempty() == true {
		return Fligh{}, errors.New("stack is empty")
	}

	index := len(s.item) - 1

	flg := s.item[index]

	return flg, nil
}

func main() {

	fmt.Println("print stack")

	s := Fligh{origin: "asd", destination: "fgk", price: 345}

	k := stack{}
	k.push(s)
	fmt.Println(k.peek())
	k.pop()
	fmt.Println(k)
	fmt.Println("over")

}
