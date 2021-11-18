package main

import (
	"errors"
	"fmt"
)

type queue struct {
	items []Flght
}

type Flght struct {
	org  string
	dest string
	cost int
}

//queue

func (q *queue) Isempty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

func (q *queue) enter(flight Flght) {

	q.items = append(q.items, flight)
	//return Flight{}
}

func (q *queue) delete() (Flght, error) {

	if len(q.items) == 0 {
		return Flght{}, errors.New("queue is empty")
	}

	//firstindex := q.items[0:len(q.items)]
	flight := q.items[1]
	q.items = q.items[1:len(q.items)]

	return flight, nil
}

func (q *queue) see() (Flght, error) {

	if len(q.items) == 0 {
		return Flght{}, errors.New("queue is empty")
	}

	flight := q.items[0]
	return flight, nil
}

func main() {

	fmt.Println("this is a queue")

	s := queue{}

	flights := []Flght{
		{org: "asd", dest: "dfg", cost: 800},
		{org: "xcv", dest: "frg", cost: 300},
		{org: "kvg", dest: "tyu", cost: 787},
	}

	s.enter(flights[2])
	fmt.Println(s.items)
	s.enter(flights[1])
	s.enter(flights[0])
	fmt.Println(s.items)
	s.delete()
	fmt.Println(s.items)
	fmt.Println(s.see())

	//fmt.Println(s)

}
