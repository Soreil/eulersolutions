package main

import (
	"fmt"
	"math"
)

func fill(n int64) []int64 {
	list := make([]int64, n-1)
	for i := int64(1); i < n-1; i++ {
		list[i-1] = i + 1
	}
	return list
}

func sieve(nPtr *[]int64) {
	n := *nPtr
	number := int64(2)
	for number*number <= int64(len(n)+2) {
		for i := number + number; i < int64(len(n)+2); i += number {
			n[i-2] = 0
		}
		number++
		for n[number-2] == 0 {
			number++
		}
	}
}

func clean(list []int64) (cleaned []int64) {
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			cleaned = append(cleaned, list[i])
		}
	}
	return
}

func factorial(n int64) int64 {
	for i := n - 1; i > 1; i-- {
		n *= i
	}
	return n
}

func digits(n int64) []rune {
	var d rune
	length := int64(math.Log10(float64(n)) + 1)
	digits := make([]rune, length)
	for i := int64(0); i < length; i++ {
		d = rune(n%10) + '0'
		n /= 10
		digits[length-i-1] = d
	}
	return digits
}

func genPermutations(n int64) []string {
	digits := digits(n)
	length := int64(math.Log10(float64(n)) + 1)
	pChan := make(chan string)
	go permute(digits, length, pChan)
	fact := factorial(length)
	retSlice := make([]string, fact)
	for i := range retSlice {
		retSlice[i] = <-pChan
	}
	return retSlice
}

func permute(s []rune, n int64, c chan string) {
	if n == 1 {
		c <- string(s)
	} else {
		for i := int64(0); i < n; i++ {
			permute(s, n-1, c)
			if n%2 == 1 {
				s[0], s[n-1] = s[n-1], s[0]
			} else {
				s[i], s[n-1] = s[n-1], s[i]
			}
		}
	}
}

func findSequence(s []int64) []int64 {
	differences := make([]int64, len(s))
	for i := 1; i < len(s); i++ {
		differences[i] = s[i] - s[0]
	}

	return differences
}

func main() {
	prime := fill(10000)
	sieve(&prime)
	prime = clean(prime)
	perms := genPermutations(1487)
	strPrime := make([]string, len(prime))
	for i, v := range prime {
		strPrime[i] = fmt.Sprint(v)
	}
	primePerms := make([]int64, 0)
	for i, v := range strPrime {
		for _, w := range perms {
			if w == v {
				primePerms = append(primePerms, prime[i])
			}
		}
	}
	PrimeSeq := findSequence(primePerms)
	fmt.Println(PrimeSeq, len(PrimeSeq))
}
