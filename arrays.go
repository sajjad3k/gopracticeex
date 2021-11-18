package arrays

import (
	"fmt"
	"strconv"
)

func main() {
	a1 := new([10]int)

	a2 := new([10]string)

	for i := 0; i < len(a1); i++ {
		a1[i] = i
	}

	for i := 0; i < len(a2); i++ {
		a2[i] = "hde" + strconv.Itoa(i)
	}

	fmt.Println(a1)
	fmt.Println(a2)

}
