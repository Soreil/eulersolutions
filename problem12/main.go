package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var primes = [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241}

// somewhere in this function there is a mistake, I get about double the count I should be getting.
// Probably shouldn't count past the middle(middle of what?) and have some bounds errors.
func factors(n int64) int64 {
	count := int64(0)
	sqrt := int64(math.Sqrt(float64(n)))
	for i := range primes {
		if n%primes[i] == 0 && primes[i] < sqrt {
			for j := int64(1); j < sqrt/primes[i]; j++ {
				if n%(primes[i]*j) == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	tmp, err := strconv.Atoi(os.Args[1])
	n := int64(tmp)
	if err != nil {
		panic(err)
	}
	count := int64(0)
	ret := int64(0)
	trianglenumber := int64(0)
	for i := int64(1); i <= n; i++ {
		trianglenumber += i
		if ret = factors(trianglenumber); ret > count {
			count = ret
			fmt.Println(count)
			if count > 500 {
				fmt.Println("result found,", trianglenumber)
			}
		}
	}
}
