package utils

import "testing"

var testCases = []struct {
	input    float64
	expected float64
}{
	{0, 1},
	{1, 1},
	{10, 3628800},
	{23, 25852016738884976640000},
	{13, 6227020800},
	{100, 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000.0},
}

func TestFactorial(t *testing.T) {
	for _, test := range testCases {
		if r := Factorial(test.input); r != test.expected {
			if !Same(r, test.expected) {
				t.Fatalf("Factorial(%g) = %g, want %g\n", test.input, r, test.expected)
			}
		}
	}
}
