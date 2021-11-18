package main

import "fmt"

func Linearsearch(items []int, n int) bool {

	for _, val := range items {

		if val == n {
			return true
		}
	}
	return false
}
func main() {

	items := []int{10, 20, 33, 44, 40, 55, 60, 77, 89, 88, 93}

	n := 11
	fmt.Println(Linearsearch(items, n))

}
