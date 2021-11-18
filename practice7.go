package main

import "fmt"

type Flights struct {
	org   string
	dest  string
	price int
}

func Getminmax(flights []Flights) (min Flights, max Flights) {

	tempmin := Flights{}
	tempmax := Flights{}

	for _, val := range flights {
		tempmin = val
		tempmax = val
		for _, k := range flights {

			if tempmin.price < k.price {
				tempmin = k
			} else if tempmax.price > k.price {

				tempmax = k
			} else {
				continue
			}

		}

	}

	return tempmin, tempmax
}

func main() {

	items := []Flights{
		{org: "asd", dest: "dfg", price: 800},
		{org: "xcv", dest: "frg", price: 300},
		{org: "kvg", dest: "tyu", price: 787},
		{org: "opl", dest: "rtg", price: 900},
		{org: "wer", dest: "yut", price: 50},
	}

	fmt.Println(Getminmax(items))

}
