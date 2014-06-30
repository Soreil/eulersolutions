package main

//Problem:
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?
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

func countprimes(n []int) int {
	count := 0
	for _, v := range n {
		if v != 0 {
			count++
		}
	}
	return count
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
	count := countprimes(sievedlist)
	fmt.Println("amount of primes", count)
}
