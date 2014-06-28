package main

import "fmt"
import "os"
import "strconv"

func summult(n int) int {
	sum := 0
	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(summult(int(n)))
}
