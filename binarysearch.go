package main

import (
	"fmt"
	"sort"
)

func Binarysearch(items []int, n int) bool {
	sort.Ints(items)

	low := 0
	high := len(items) - 1

	for low <= high {
		med := (low + high) / 2

		if items[med] < n {
			low = med + 1 // return true
		} else {
			high = med - 1
		}
	}

	if low == len(items) || items[low] != n {
		return false
	}
	return true
}
func main() {
	items := []int{10, 20, 33, 44, 40, 55, 60, 99, 45, 66, 21}

	//sort.Ints(items)
	n := 33
	fmt.Println(Binarysearch(items, n))
	//fmt.Println(items)

}
