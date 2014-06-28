package main

import "fmt"
import "os"
import "strconv"

func smallestMultiple(n int) (sum int) {
	sum = 1
	for i := 2; i <= n; i++ {
		prevsum := sum
		for sum%i != 0 {
			sum += prevsum
		}
	}
	return
}

func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(smallestMultiple(int(n)))
}
