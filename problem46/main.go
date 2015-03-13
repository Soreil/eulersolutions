package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Soreil/eulersolutions/utils"
)

/*
Sieve a number range.
Save both the list of primes and the list of composites.
Make a list of doubled squares up to the highest prime found.

Subtract the doubled square from a composite number, check if the result is prime. if not try another square.
If no double square subtraction produces a prime the conjecture is false.
*/
var primes []int
var composites []int
var dSquares []int

func main() {

	lim := 10000

	primes = utils.Sieve(lim)
	composites = sieveComposites(primes)
	dSquares = genDSquares(lim)

	for _, c := range composites {
		if c%2 == 0 {
			continue
		}
		disproven := true
		for _, d := range dSquares {

			if isPrime(c - d) {
				disproven = false
			}
			if d >= c {
				if disproven == true {
					fmt.Println("Conjecture false", c)
					os.Exit(0)
				} else {
					break
				}
			}
		}
	}
}

func sieveComposites(p []int) (c []int) {
	for i := 0; i <= p[len(p)-1]; i++ {
		c = append(c, i)
	}
	//offset is compensation for the shrinking array so indexing works
	var offset int
	for _, v := range p {
		c = append(c[:v-offset], c[v-offset+1:]...)
		offset++
	}
	return c[2:]
}

func genDSquares(lim int) []int {
	var squares []int
	sqLim := int(math.Sqrt(float64(lim)))
	for i := 1; i <= sqLim; i++ {
		squares = append(squares, 2*(i*i))
	}
	return squares
}

func isPrime(n int) bool {
	for _, v := range primes {
		if n == v {
			return true
		}
		if v > n {
			return false
		}
	}
	return false
}
