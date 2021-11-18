package main

import "fmt"

type Developer struct {
	name string
}

func unique(dev []Developer) []string {

	var unq []string
	tmp := false

	for _, val := range dev {
		for _, kal := range unq {
			if val.name == kal {
				tmp = true
			}
		}
		if !tmp {
			unq = append(unq, val.name)
		}
	}

	return unq
}

func main() {

	dev := []Developer{
		{name: "Elliot"},
		{name: "Alan"},
		{name: "Jennifer"},
		{name: "Graham"},
		//	{name: "Elliot"},
		{name: "Paul"},
		//{name: "Alan"},
		//{name: "Graham"},
		//{name: "Alan"},
	}

	out := unique(dev)

	fmt.Println(out)

}
