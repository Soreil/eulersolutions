package main

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n, count int) int {
	count++
	if n%2 == 0 {
		return collatz(n/2, count)
	} else if n != 1 {
		return collatz((3*n)+1, count)
	}
	return count
}
func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	biggest := 0
	tmp := 0
	for i := 1; i < n; i++ {
		if tmp = collatz(i, 0); tmp > biggest {
			biggest = tmp
			fmt.Println(biggest, i)
		}
	}
}
