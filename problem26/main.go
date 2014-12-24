package main

import "fmt"

var lim int64 = 50

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
	for i := int64(2); i < lim; i++ {
		//Coprimes of 10(multiples of 2 and 5) can't be cyclic
		if coprime10(i) {
			continue
		}
		var j int64
		var tmp int64
		for tmp, j = i, 1; gcd(tmp, 10) != 1; j++ {
			tmp *= i
		}
		if j == 1 {
			fmt.Println(i, " ", i-1)
		} else {
			fmt.Println(i, " ", j-1, "times something")
		}
	}
}
