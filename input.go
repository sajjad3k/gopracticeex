package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter num")
	k, _ := reader.ReadString('\n')

	// remove newline
	str1 := strings.Replace(k, "\n", "", -1)

	// convert string variable to int variable
	num, _ := strconv.Atoi(str1)

	if num >= 1 && num <= 10 {
		fmt.Println("yes")
	}

}
