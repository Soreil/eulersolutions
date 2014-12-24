package main

import "fmt"

/*
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?
*/

func main() {
	perms(200, 1, 2, 5)
}

func perms(lim int, vals ...int) int {
	count := 0
	for i := range vals {
		if i == 0 {
			count += 1
			continue
		}
		//2 or 2.5
		ratio := float64(vals[i] / vals[i-1])
		//multiply count by lim / ratio
		count += count * int(float64(lim)/ratio)
		fmt.Println(count)
	}
	return count
}
