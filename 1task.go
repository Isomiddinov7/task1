package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	fmt.Println((a))
}

func computer(a int) (res string) {
	computers := " компьютеров"
	computer := " компьютер"

	if a > 1 {
		res = strconv.Itoa(a) + computers
	} else {
		res = strconv.Itoa(a) + computer
	}

	return res
}
