package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib() func() int {
	a, b := 1, 2
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	i, sum := 0, 0
	f := fib()
	for {
		i = f()
		if i%2 == 0 {
			if i <= int(n) {
				sum += i
			} else {
				break
			}
		}
	}
	fmt.Println(sum)
}
