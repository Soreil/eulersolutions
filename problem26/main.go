package main

import "fmt"

func coprime10(n int64) bool {
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

//Calculates the greatest commmon divisor of two numbers
func gcd(x, y int64) int64 {
	var min int64
	if x < y {
		min = x
	} else {
		min = y
	}
	for i := min; i > 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1

}

func main() {

	var lim int64 = 1000

	for d := int64(2); d < lim; d++ {
		//Coprimes of 10(multiples of 2 and 5) can't be cyclic
		if coprime10(d) {
			continue
		}
		var i int64
		var tmp int64
		for tmp, i = d, 1; gcd(tmp, 10) != 1; i++ {
			tmp *= d
		}
		if i == 1 {
			fmt.Println(d, " ", d-1)
		} else {
			fmt.Println(d, " ", i-1, "times something")
		}
	}
}
