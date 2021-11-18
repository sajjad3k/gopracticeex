package main

import (
	"errors"
	"fmt"
)

func Checkpermutation(s1, s2 string) (bool, error) {

	var per int = 0

	if len(s1) != len(s2) {
		return false, errors.New("string length unmatched")
	}

	for _, v := range s1 {
		for _, d := range s2 {
			//if strings.Contains()
			if v == d {
				per++
			} else {
				continue
			}
		}
	}

	if per == len(s1) {
		return true, nil
	} else {
		return false, errors.New("no permutatio")
	}
}

func main() {

	s1 := "abcdef"
	s2 := "fedcba"

	fmt.Println(Checkpermutation(s1, s2))

}
