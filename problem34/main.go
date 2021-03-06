package main

import (
	"fmt"

	"github.com/Soreil/eulersolutions/utils"
)

func main() {
	//The number 9.999.999 is the largest amount of 9s
	//we can have before the result is ALWAYS smaller than the original number
	//Make a slice with the factorials so we don't duplicate computation
	ffact := make([]int, 10)
	for i := range ffact {
		ffact[i] = utils.Factorial(i)
	}
	lim := ffact[9] * 7
	var sum int
	var tmpSum int
	//Do not count 1 and 2
	for i := 3; i < lim; i++ {
		tmpSum = 0
		for n := i; n > 0; n /= 10 {
			tmpSum += ffact[n%10]
		}
		if tmpSum == i {
			sum += tmpSum
		}
	}
	fmt.Printf("Sum of magic numbers:%d\n", sum)
}
