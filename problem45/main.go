package main

import (
	"fmt"

	"gopkg.in/fatih/set.v0"
)

func triangle(lim int) *set.Set {
	s := set.New()
	for i := 1; i < lim; i++ {
		s.Add(i * (i + 1) / 2)
	}
	return s
}

func pent(lim int) *set.Set {
	s := set.New()
	for i := 1; i < lim; i++ {
		s.Add(i * (3*i - 1) / 2)
	}
	return s
}

func hex(lim int) *set.Set {
	s := set.New()
	for i := 1; i < lim; i++ {
		s.Add(i * (2*i - 1))
	}
	return s
}

func main() {
	t := triangle(100000)
	p := pent(100000)
	h := hex(100000)
	common := set.Intersection(t, p, h)
	fmt.Println(common)
}
