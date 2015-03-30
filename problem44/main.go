package main

import (
	"fmt"
	"math"
)

func main() {
	//Pn = n(3n-1)/2
	pen := genPen(100000)
	for pj := range pen {
		for pk := range pen[pj:] {
			if isPent(pen[pj]+pen[pk]) && isPent(pen[pk]-pen[pj]) {
				fmt.Printf("D:%g\n", math.Abs(float64(pen[pk]-pen[pj])))
				fmt.Printf("Pj:%d Pk:%d\n", pj, pk)
			}
		}
	}
}

func genPen(n int) []int {
	var p []int

	for i := 1; i <= n; i++ {
		p = append(p, i*(3*i-1)/2)
	}
	return p
}

//In case this produces an integer it's a pent
func isPent(x int) bool {
	if x == 0 {
		return false
	}
	tmp := math.Sqrt(float64(24*x+1)) + 1/6
	return tmp == math.Floor(tmp)
}
