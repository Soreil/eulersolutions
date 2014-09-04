package main

import (
	"fmt"
	"math/big"
)

func fibonacci() func() string {
	a := new(big.Int)
	b := new(big.Int)
	tmp := new(big.Int)
	a.SetInt64(0)
	b.SetInt64(1)
	return func() string {
		tmp = tmp.Set(a)
		a = a.Set(b)
		b = b.Add(tmp, b)
		return fmt.Sprint(a)
	}
}

func main() {
	f := fibonacci()
	var str string
	var i int
	for i = 0; len(str) < 1000; i++ {
		str = f()

	}
	fmt.Println(str, i)
}
