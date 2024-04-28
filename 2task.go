package main

import "fmt"

func main() {
	req := []int{12, 24, 8}
	fmt.Println(AllDiv(req))
}

func AllDiv(a []int) []int {
	var (
		count  int
		k      = []int{}
		length = len(a)
		b      = 2
		min    = a[0]
	)

	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}

	for j := 0; j < length; j++ {
		if a[j]%b == 0 {
			count++
		} else {
			b++
			j = -1
		}
		if count == length {
			k = append(k, b)
			count = 0
			b++
			j = -1
			if b == min {
				break
			}
		}
	}

	return k
}
