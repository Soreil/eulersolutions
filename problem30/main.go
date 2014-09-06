package main

import (
	"fmt"
	"math"
)

func main() {
	total := float64(0)
	for i := 2; i < 1000000; i++ {
		sum := float64(0)
		str := fmt.Sprintf("%05d", i)
		for _, v := range str {
			sum += math.Pow(float64(v-'0'), 5)
		}
		if sum == float64(i) {
			fmt.Println(sum)
			total += sum
		}
	}
	fmt.Println(total)
}
