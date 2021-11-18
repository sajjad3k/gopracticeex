package main

import (
	"fmt"
	"sort"
)

type Flight struct {
	origin      string
	destn       string
	ticketprice int
}

func Sortflights(flights []Flight) []Flight {

	var sorted []Flight = flights

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ticketprice < sorted[j].ticketprice
	})

	/*temp := Flight{ticketprice: 0}

	for _, fl := range flights {
		//temp.ticketprice = fl.ticketprice
		if temp.ticketprice < fl.ticketprice {
			temp.origin = fl.origin
			temp.destn = fl.destn
			temp.ticketprice = fl.ticketprice
			sorted = append(sorted, temp)
		}

		temp = Flight{}

	} */
	return sorted
}

//sorting example
func main() {

	flights := []Flight{
		{origin: "asd", destn: "dfg", ticketprice: 800},
		{origin: "xcv", destn: "frg", ticketprice: 300},
		{origin: "kvg", destn: "tyu", ticketprice: 787},
	}

	out := Sortflights(flights)
	fmt.Println(out)
}
