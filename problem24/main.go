package main

import "fmt"

func lex(shift []int, size int) (lex []int) {
	base := make([]int, size)
	lex = make([]int, size)
	for i := 0; i < size; i++ {
		base[i] = i
	}
	for i, v := range shift {
		lex[i] = base[v]
		base = append(base[:v], base[v+1:]...)
	}
	return
}

func main() {
	num := 1000000 - 1
	size := 10
	permutation := make([]int, size)
	mods := []int{362880, 40320, 5040, 720, 120, 24, 6, 2, 1}
	for i := 0; i < size-1; i++ {
		permutation[i] = num / mods[i]
		num %= mods[i]
	}
	result := lex(permutation, size)
	fmt.Println(result)
}
