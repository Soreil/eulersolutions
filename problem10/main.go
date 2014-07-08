package main

//Problem:
//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//Find the sum of all the primes below two million.

import "fmt"
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

func sumprimes(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func main() {
	var n int
	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		panic("Enter a number")
	}

	list := make([]int, n)

	fill(list)
	sievedlist := sieve(list)
	sum := sumprimes(sievedlist)

	fmt.Printf("Sum of primes below %d: %d\n", n, sum)
}
