package main

//Problem:
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?

//Solution:
//Use the Sieve of Arasthosis to find the actual primes
//Use n (log n + log log n) to set the Sieve size to just over where the nth prime should be
//Count until we hit the nth prime

import "fmt"
import "math"
import "os"
import "strconv"

func sieve(n []int) []int {
	number := 2
	for number*number <= len(n)+2 {
		for i := number + number; i < len(n)+2; i += number {
			n[i-2] = 0
		}
		number++
		for n[number-2] == 0 {
			number++
		}
	}
	return n
}

func fill(n []int) []int {
	for i := 1; i <= len(n); i++ {
		n[i-1] = i + 1
	}
	return n
}

func countprimes(n []int, m int) (int, int) {
	count := 0
	nth := 0
	for _, v := range n {
		if v != 0 {
			count++
			if count == m {
				nth = v
			}
		}
	}
	return count, nth
}

func approx(n float64) int {
	return int(n * (math.Log(n) + math.Log(math.Log(n))))
}

func main() {
	var n int
	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		panic("Enter a number")
	}

	estimate := float64(approx(float64(n)))
	list := make([]int, int(estimate))
	fill(list)

	sievedlist := sieve(list)

	count, nth := countprimes(sievedlist, n)
	fmt.Println("amount of primes", count)
	if nth != 0 {
		fmt.Printf("%dth prime: %d\n", n, nth)
	} else {
		fmt.Printf("Failed to find the %dth prime!\n", n)
	}
}
