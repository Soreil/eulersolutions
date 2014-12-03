package main

import (
	"fmt"

	"github.com/Soreil/eulersolutions/utils"
)

var alim int = 1000
var blim int = 1000

var primes []int = utils.Sieve(100000)

func main() {
	a, b := -999, -999
	// test values that produce 40 primes, go above 1000
	a = 1
	b = 41
	length, longest := 0, 0
	//Figure out a logical path for decrementing a and b
	//for math.Abs(float64(a)) < float64(alim) && math.Abs(float64(b)) < float64(blim) {
	length = 0
	//no idea how big n will get, shouldn't need a limit
	for n := 0; n < 1000; n++ {
		if isPrime(genPrime(a, b, n)) {
			length++
		} else {
			fmt.Println("Not a winner", a, b, n)
			fmt.Println(genPrime(a, b, n))
			break
		}
	}
	if length > longest {
		longest = length
		fmt.Println(longest)
	}
	//}
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
