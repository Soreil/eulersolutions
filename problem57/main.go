package main

import "fmt"
import "math/big"

func main() {
	fmt.Println(sqrtConverge(3, 2, 1000))
}

func sqrtConverge(n, d, lim int64) int64 {
	var count int64
	ratio := big.NewRat(n, d)
	tmp := big.NewRat(n, d)
	for ; lim > 0; lim-- {
		//n, d = n + (2 * d), n + d
		tmp.Num().Add(tmp.Num(), tmp.Denom())
		tmp.Num().Add(tmp.Num(), tmp.Denom())

		tmp.Denom().Add(tmp.Denom(), ratio.Num())

		ratio.SetFrac(tmp.Num(), tmp.Denom())

		if len(fmt.Sprint(ratio.Num().String())) > len(fmt.Sprint(ratio.Denom().String())) {
			count++
		}
	}
	return count
}
