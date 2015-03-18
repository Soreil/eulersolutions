package main

import "testing"

var testCases = []struct {
	hand     hand
	expected bool
}{
	{
		hand{
			card{king, spades},
			card{king, spades},
			card{king, spades},
			card{king, spades},
			card{king, spades},
		},
		false,
	},
	{
		hand{
			card{9, spades},
			card{jack, spades},
			card{queen, spades},
			card{king, spades},
			card{ace, spades},
		},
		false,
	},
	{
		hand{
			card{10, spades},
			card{jack, spades},
			card{queen, hearts},
			card{king, spades},
			card{ace, spades},
		},
		false,
	},
	{
		hand{
			card{10, spades},
			card{jack, spades},
			card{queen, spades},
			card{king, spades},
			card{ace, spades},
		},
		true,
	},
}

func Test(t *testing.T) {
	for _, test := range testCases {
		if r := test.hand.isRoyalFlush(); r != test.expected {
			t.Fatalf("isRoyalFlush(%v):%v, expected:%v\n", test.hand, test.expected, test.expected)
		}
	}
}
