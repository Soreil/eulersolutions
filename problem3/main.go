package main

import "fmt"
import "os"
import "strconv"

func factor(n int) []int {
	s := make([]int, 0)
	d := 2
	for n > 1 {
		for n%d == 0 {
			s = append(s, d)
			n /= d
		}
		d++
	}
	return s
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	s := factor(int(n))
	fmt.Println(s)
}
