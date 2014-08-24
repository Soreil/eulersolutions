package main

import (
	"fmt"
	"math"
	"os"
)

func pyTripple(m, n int) (a, b, c int) {
	a = 2 * m * n
	b = m*m - n*n
	c = m*m + n*n
	return
}

func main() {
	m := 2
	n := 1
	a, b, c := 1, 1, 1
	sum := a + b + c
	target := 1000
	limit := int(math.Sqrt(float64(target)))
	for m = 2; target%sum != 0 && m < limit; m++ {
		a, b, c = pyTripple(m, n)
		sum = a + b + c
		fmt.Println(a, b, c, sum)
		if target%sum == 0 {
			fmt.Println(a*target/sum, b*target/sum, c*target/sum)
			os.Exit(0)
		} else {
			fmt.Println(target % sum)
			for n = 1; n < m; n++ {
				a, b, c = pyTripple(m, n)
				sum = a + b + c
				fmt.Println(a, b, c, sum)
				if target%sum == 0 {
					fmt.Println(a*target/sum, b*target/sum, c*target/sum)
					os.Exit(0)
				} else {
					fmt.Println(target % sum)
				}

			}
		}
	}
}
