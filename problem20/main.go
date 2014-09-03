package main

import (
	"fmt"
	"math/big"
)

func sumDigits(digits string) (sum int) {
	for _, r := range digits {
		sum += int(r - '0')
	}
	return
}

func main() {
	result := new(big.Int)
	result = result.MulRange(1, 100)
	fmt.Println(sumDigits(fmt.Sprint(result)))
}
