package main

import "testing"

var testCase = []struct {
	input    int
	expected bool
}{
	{
		55,
		true,
	},
	{
		56,
		false,
	},
	{
		45,
		true,
	},
}

func Test(t *testing.T) {
	for _, test := range testCase {
		if r := isTriangular(test.input); r != test.expected {
			t.Fatalf("isTriangular(%d) = %t, expected:%t\n", test.input, r, test.expected)
		}
	}
}
