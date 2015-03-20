package main

import (
	"testing"

	"github.com/Soreil/eulersolutions/utils"
)

type test []struct {
	a        int
	b        int
	expected int
}

var testCases = test{
	{
		640,
		480,
		160,
	},
	{
		100,
		160,
		20,
	},
	{
		100,
		150,
		50,
	},
	{
		2,
		8,
		2,
	},
	{
		2,
		7,
		1,
	},
	{
		11,
		12,
		1,
	},
}

func TestGcd(t *testing.T) {
	p := utils.Sieve(10000)
	for _, v := range testCases {
		if r := gcd(v.a, v.b, p); r != v.expected {
			t.Fatalf("gcd(%d,%d) = %d, expected:%d\n", v.a, v.b, r, v.expected)
		}
	}
}
