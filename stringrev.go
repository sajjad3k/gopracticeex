package main

import (
	"fmt"
	"strings"
)

func main() {

	inp := "a,bcd fgt$love#"

	tem := []string{}
	
	spec := make(map[int]string)

	out := []string{}

	for i, val := range inp {

		if (rune(val) >= 'a' && rune(val) <= 'z') || (rune(val) >= 'A' && rune(val) <= 'Z') {
			//tem[string(val)] = i
			tem = append(tem, string(val))
		} else {
			spec[i] = string(val)
		}
	}

	
	n := len(tem) - 1

	for i, v := range inp {
		if (rune(v) >= 'a' && rune(v) <= 'z') || (rune(v) >= 'A' && rune(v) <= 'Z') {
			out = append(out, tem[n])
			n--
		}
		if string(v) == spec[i] {

			out = append(out, spec[i])
		}
	}

	fmt.Println(inp)
	fmt.Println(strings.Join(out, ""))

}
