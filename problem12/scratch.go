package main

import "fmt"
import "math"
import "os"
import "strconv"

func get_trangulars(n int) []int {
	list := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		list[i] = list[i-1] + i
	}
	return list
}

func factor(n int) []int {
	list := make([]int, 0)
	list = append(list, num)
	return list
}
func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	total := get_trangulars(int(n))
	factors := factor(total[n])
	fmt.Println(factors)
}
