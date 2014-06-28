package main

import "fmt"
import "math"
import "os"
import "strconv"

var highestp int = 0

func fetch(n int) func() []int {
	x := int(math.Pow10(n))
	return func() []int {
		if x > 1 {
			x--
			s := make([]int, x)
			for i := x; i > 0; i-- {
				s[i-1] = x * i
			}
			return s
		}
		return nil
	}
}

func reversestr(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
func checkpalindrome(n int) bool {
	palinstr := strconv.FormatInt(int64(n), 10)
	return palinstr == reversestr(palinstr)
}

func ispalindrome(n int) bool {
	if checkpalindrome(n) && n > highestp {
		highestp = n
		return true
	}
	return false
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	f := fetch(int(n))
	for r := f(); r != nil; r = f() {
		for _, v := range r {
			ispalindrome(v)
		}
	}
	if highestp == 0 {
		fmt.Println("Failed to find a palindrome")
	} else {
		fmt.Printf("Highest palindrome:%d\n", highestp)
	}
}
