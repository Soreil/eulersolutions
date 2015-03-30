package main

import (
	"testing"

	"github.com/Soreil/eulersolutions/utils"
)

type input struct {
	n float64
	c float64
}

var testCases = []struct {
	input    input
	expected float64
}{
	{input{n: 5, c: 3}, 10},
	{input{n: 3, c: 4}, -1},
	{input{n: 23, c: 10}, 1144066},
	{input{n: 100, c: 50}, 100891344545564193334812497256},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		if r := ncr(test.input.n, test.input.c); r != test.expected {
			if !utils.Same(r, test.expected) {
				t.Fatalf("ncr(%g,%g) = %g, want %g\n", test.input.n, test.input.c, r, test.expected)
			}
		}
	}
}
