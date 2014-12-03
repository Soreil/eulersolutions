package main

import (
	"fmt"
	"math"

	"github.com/Soreil/eulersolutions/utils"
)

var primes []int = utils.Sieve(10000)

func main() {
	var alim int = 1000
	var blim int = 1000
	longest := 0
	//Figure out a logical path for decrementing a and b
	for a := -1000; math.Abs(float64(a)) <= float64(alim); a++ {
		for b := 0; b < blim; b++ {
			n := 0
			for {
				if isPrime(genPrime(a, b, n)) {
					n++
				} else {
					break
				}
			}
			if n > longest {
				longest = n
				fmt.Println(a, b, n)
			}
		}
	}
}

func genPrime(a, b, n int) int {
	return n*n + a*n + b
}

func isPrime(n int) bool {
	for _, v := range primes {
		if v == n {
			return true
		} else if v > n {
			return false
		}
	}
	return false
}
