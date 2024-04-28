package main

import "fmt"

func main() {
	var min, max, count int
	var k = []int{}
	fmt.Scan(&min, &max)

	for i := min; i <= max; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count < 2 {
			k = append(k, i)

		}
		count = 0
	}

	fmt.Println(k)
}
