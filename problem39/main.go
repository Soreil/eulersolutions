package main

import (
	"fmt"
	"sort"
)

type tripples [][3]int

func (t tripples) Len() int {
	return len(t)
}

func (t tripples) Less(i, j int) bool {
	return t[i][0]+t[i][1]+t[i][2] < t[j][0]+t[j][1]+t[j][2]
}

func (t tripples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

//Returns a primitive pythagorean tripple
func pyTripple(m, n int) (a, b, c int) {
	a = m*m - n*n
	b = 2 * m * n
	c = m*m + n*n
	return
}

//Numbers are coprime when their greatest common divisor is one
func isCoPrime(a, b int) bool {
	return gcd(a, b) == 1
}

func perimeter(sides [3]int) int {
	return sides[0] + sides[1] + sides[2]
}

//Calculates the greatest commmon divisor of two numbers
func gcd(x, y int) int {
	var min int
	if x < y {
		min = x
	} else {
		min = y
	}
	for i := min; i > 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	var tripples tripples

	pMax := 1000 //The search space for biggest amount of roots
	pBiggest := 0
	pBiggestCount := 0

	//The limit here is arbitrary
	//TODO:(sjon) figure out whether no smaller circumference primitives can be found at some point. They do not appear to be sorted.
	for m := 2; m < 30; m++ {
		for n := 1; n < m; n++ {
			//if m-n is uneven it can't possibly be a primitive tripple
			if (m-n)%2 == 0 {
				//We need coprimality
				if isCoPrime(m, n) {
					a, b, c := pyTripple(m, n)
					//if both m and n are uneven we have a doubled tripple
					if m%2 == 1 && n%2 == 1 {
						a /= 2
						b /= 2
						c /= 2
					}
					tripples = append(tripples, [3]int{a, b, c})
				}
			}
		}
	}
	sort.Sort(tripples)
	for p := 2; p < pMax; p++ {
		var pCount int
		for i := range tripples {
			if perimeter(tripples[i]) > p {
				break
			}
			if p%perimeter(tripples[i]) == 0 {
				pCount++
			}
		}
		if pCount > pBiggestCount {
			pBiggestCount = pCount
			pBiggest = p
		}
	}
	fmt.Println(pBiggestCount)
	fmt.Println(pBiggest)
}
