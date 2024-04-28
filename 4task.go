package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	for i := 0; i <= a; i++ {
		for j := 1; j <= a; j++ {
			if i == 0 {
				fmt.Printf("  " + strconv.Itoa(j))
			} else {
				fmt.Printf(" " + strconv.Itoa(i*j))
			}
		}

		fmt.Printf("\n" + strconv.Itoa(i+1))

	}
}
