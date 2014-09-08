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

func digits(n int64) []int64 {
	var d int64
	length := int64(math.Log10(float64(n)) + 1)
	digits := make([]int64, length)
	for i := int64(0); i < length; i++ {
		d = n % 10
		n /= 10
		digits[length-i-1] = d
	}
	return digits
}

func permutations(n int64) [][]int64 {
	length := int64(math.Log10(float64(n)) + 1)
	fact := factorial(length)
	perms := make([][]int64, fact)
	digits := digits(n)
	perms[0] = digits
	for i := 1; i < len(perms); i++ {
		perms[i] = digits
	}

	return perms
}

func main() {
	list := fill(10000)
	sieve(&list)
	list = clean(list)
	fmt.Println(permutations(1478))
}
