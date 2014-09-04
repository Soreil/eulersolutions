package main

import (
	"fmt"
	"math/big"
)

func suffixDigits(n int, digits string) (s string) {
	s = digits[len(digits)-n:]
	return
}

func main() {
	result := new(big.Int)
	tmp := new(big.Int)
	for i := int64(1); i <= 1000; i++ {
		tmp = tmp.SetInt64(i)
		tmp = tmp.Exp(tmp, tmp, nil)
		result.Add(result, tmp)
	}
	fmt.Println(suffixDigits(10, fmt.Sprint(result)))
}
