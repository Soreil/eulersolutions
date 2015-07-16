package main

import (
	"fmt"
	"log"

	"github.com/Soreil/eulersolutions/utils"
)

func coprime10(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%5 == 0 {
		n /= 5
	}
	if n == 1 {
		return true
	}
	return false
}

//
////Calculates the greatest commmon divisor of two numbers
//func gcd(x, y int64) int64 {
//	var min int64
//	if x < y {
//		min = x
//	} else {
//		min = y
//	}
//	for i := min; i > 1; i-- {
//		if x%i == 0 && y%i == 0 {
//			return i
//		}
//	}
//	return 1
//
//}

func isCyclic(p int) bool {
	b := 10
	t := 0
	r := 1
	n := 0
	for {
		t++
		x := r * b
		d := x / p
		r = x % p
		n = n*b + d
		if r == 1 {
			break
		}
		log.Println(n)
	}
	return t == p-1
}

func main() {
	var lim int = 1000
	primes := utils.Sieve(lim)
	var max int
	for d := 2; d <= lim; d++ {
		//Coprimes of 10(multiples of 2 and 5) can't be cyclic
		if coprime10(d) {
			continue
		}
		if utils.IsPrime(d, primes) {
			//	fmt.Println(d, isCyclic(d))
			if isCyclic(d) {
				max = d
			}
		}
	}
	fmt.Println("Result:", max)
}
