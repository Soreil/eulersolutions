package main

import (
	"log"

	"github.com/Soreil/eulersolutions/utils"
)

/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?

A:
Primes other than 2 and 5 need to end with 1,3,7,9.
this being because any multiple of 5(a prime) ends in 5 and the other numbers are even.
Therefore a rotation can not have digits other than 1,3,7,9.
*/

func main() {

	primes := utils.Sieve(100)
	log.Println(primes)
	filtered := filter(primes)
	//Generate the primes up to one million
	//Filter out all primes containing numbers that can't be part of a cyclical prime
	//Check if every rotation is in the list
	//If every rotation is in the list remove all from the list and add to the cyclical primes count
	log.Println(filtered)
}

func filter(primes []int) []int {
	var out []int
	for i, v := range primes {
		if v < 10 {
			out = append(out, v)
			continue
		}
		for !is024568(v % 10) {
			v /= 10
			if v == 0 {
				out = append(out, primes[i])
				break
			}
		}
	}
	return out
}

func is024568(c int) bool {
	return c == 0 || c == 2 || c == 4 || c == 5 || c == 6 || c == 8
}
