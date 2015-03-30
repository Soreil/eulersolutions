package main

import (
	"fmt"

	"github.com/Soreil/eulersolutions/utils"
)

func main() {
	var mills int
	for n := float64(0); n <= 100; n++ {
		for r := float64(0); r < n; r++ {
			if ncr(n, r) > 1000000 {
				mills++
			}
		}
	}
	fmt.Println(mills)
}

func ncr(n, r float64) float64 {
	if r > n {
		return -1
	}
	return (utils.Factorial(n) / (utils.Factorial(r) * utils.Factorial(n-r)))
}
