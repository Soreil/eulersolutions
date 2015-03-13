package main

import "fmt"

func isPalin(n int) bool {
	return isDecPalin(n) && isBinPalin(n)
}

func isDecPalin(n int) bool {
	return fmt.Sprint(n) == reverseString(fmt.Sprint(n))
}

func isBinPalin(n int) bool {
	return fmt.Sprintf("%b", n) == reverseString(fmt.Sprintf("%b", n))
}

func reverseString(s string) string {
	var rev string
	for i := range s {
		rev += string(s[len(s)-1-i])
	}
	return rev
}

func main() {
	var sum int
	for i := 1; i < 1000000; i += 2 {
		if isPalin(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
