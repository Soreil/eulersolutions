package main

import (
	"fmt"

	"github.com/Soreil/eulersolutions/utils"
)

var primes []int

func main() {
	primes = utils.Sieve(1000000)

	var sum []int
	var prime int

	for _, v := range primes {
		if tmpSum := consPrimeSum(v, 0); len(tmpSum) > len(sum) {
			sum = tmpSum
			prime = v
		}
	}
	fmt.Println("Prime:", prime)
}

func consPrimeSum(p int, count int) []int {
	if count > 10 {
		return []int{}
	}
	var sum int
	for i := range primes {
		if i+count >= len(primes) {
			break
		}
		sum += primes[i+count]
		if sum > p {
			return consPrimeSum(p, count+1)
		} else if sum == p {
			return primes[count : i+count+1]
		}
	}
	return []int{}
}
