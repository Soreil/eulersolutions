package main

import "fmt"

func spiralDiagonals(n int64) (sum int64) {
	sum = 1
	for n > 1 {
		sum += n*n*4 - ((n - 1) * 6) // The 6 is 0+1+2+3
		n -= 2
	}
	return
}

func main() {
	n := int64(1001)
	numbers := spiralDiagonals(n)
	fmt.Println(numbers)
}
