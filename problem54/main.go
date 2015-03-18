package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
)

type suit int
type value int

type card struct {
	value
	suit
}

type hand []card

func (h hand) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hand) Len() int {
	return len(h)
}

const (
	jack value = 11 + iota
	queen
	king
	ace
)

const (
	hearts suit = iota
	clubs
	spades
	diamonds
)

func main() {
	a := hand{{jack, hearts}, {3, clubs}, {4, hearts}, {5, clubs}, {10, diamonds}}
	b := hand{{jack, hearts}, {10, clubs}, {4, hearts}, {5, clubs}, {10, diamonds}}
	won, err := a.isWin(b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if won {
		fmt.Println("A won!")
	} else {
		fmt.Println("B won!")
	}
}

//return true if A has a higher hand
func (a hand) isWin(b hand) (bool, error) {
	if len(a) != 5 || len(b) != 5 {
		return false, errors.New("Not a 5 card hand")
	}

	sort.Sort(a)
	sort.Sort(b)

	//Check if a has a pair
	//r := sort.Search(len(a)-1, func(i int) bool { return a[i].value == a[i+1].value })

	//Check if b has a pair
	//r = sort.Search(len(b)-1, func(i int) bool { return b[i].value == b[i+1].value })
	return true, nil
}

func (h hand) isRoyalFlush() bool {
	//DeepEqual is needed since this is not a type I want to write a stringer and wrapper for
	if reflect.DeepEqual([]value{10, jack, queen, king, ace}, []value{h[0].value, h[1].value, h[2].value, h[3].value, h[4].value}) {
		//Pick the first suit we find
		suit := h[0].suit
		if h[1].suit == suit && h[2].suit == suit && h[3].suit == suit && h[4].suit == suit {
			return true
		}
	}
	return false
}
func (h hand) isStraightFlush() bool {
	//Check if all cards are in repeating order, Ace-2-3-4-5 is unhandled because it wasn't listed
	if h[1].value == h[0].value+1 && h[2].value == h[1].value+1 && h[3].value == h[2].value+1 && h[4].value == h[3].value+1 {
		//Pick the first suit we find
		suit := h[0].suit
		if h[1].suit == suit && h[2].suit == suit && h[3].suit == suit && h[4].suit == suit {
			return true
		}
	}
	return false
}

func (h hand) fourOfAKind() bool {
	if h[0].value == h[1].value && h[1].value == h[2].value && h[2].value == h[3].value && h[3].value == h[4].value {
		return true
	}
	return false
}
