package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func bigPower(exponent int) string {
	sum := big.NewInt(0)
	sum.SetBit(sum, exponent, 1)
	return fmt.Sprint(sum)
}

func sumDigits(digits string) (sum int) {
	for _, r := range digits {
		sum += int(r - '0')
	}
	return
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(sumDigits(bigPower(n)))
}
