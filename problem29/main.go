package main

import (
	"fmt"
	"math/big"

	"gopkg.in/fatih/set.v0"
)

func main() {
	lim := 100
	combinations := set.New()
	iBig := new(big.Int)
	jBig := new(big.Int)
	b := new(big.Int)
	for i := 2; i <= lim; i++ {
		iBig.SetInt64(int64(i))
		for j := 2; j <= lim; j++ {
			jBig.SetInt64(int64(j))
			b.Exp(iBig, jBig, nil)
			combinations.Add(b.String())
		}
	}
	fmt.Println(combinations.Size())
}
