package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Problem 17
//
//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//
//note: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

func sumOfLetters(n int) (sum int64) {
	for i := 1; i <= n; i++ {
		sum += int64(len(word(strings.Split(fmt.Sprint(i), ""))))
	}
	return
}
func word(n []string) string {
	var str []string

	if len(n) >= 2 {
		switch n[len(n)-2] {
		case "2":
			str = append(str, "twenty")
		case "3":
			str = append(str, "thirty")
		case "4":
			str = append(str, "forty")
		case "5":
			str = append(str, "fifty")
		case "6":
			str = append(str, "sixty")
		case "7":
			str = append(str, "seventy")
		case "8":
			str = append(str, "eighty")
		case "9":
			str = append(str, "ninety")
		case "1":
			switch n[len(n)-1] {
			case "0":
				str = append(str, "ten")
			case "1":
				str = append(str, "eleven")
			case "2":
				str = append(str, "twelve")
			case "3":
				str = append(str, "thirteen")
			case "4":
				str = append(str, "fourteen")
			case "5":
				str = append(str, "fifteen")
			case "6":
				str = append(str, "sixteen")
			case "7":
				str = append(str, "seventeen")
			case "8":
				str = append(str, "eighteen")
			case "9":
				str = append(str, "nineteen")
			}

		}
	}
	if n[len(n)-2] != "1" {
		switch n[len(n)-1] {
		case "1":
			str = append(str, "one")
		case "2":
			str = append(str, "two")
		case "3":
			str = append(str, "three")
		case "4":
			str = append(str, "four")
		case "5":
			str = append(str, "five")
		case "6":
			str = append(str, "six")
		case "7":
			str = append(str, "seven")
		case "8":
			str = append(str, "eight")
		case "9":
			str = append(str, "nine")
		}
	}
	fmt.Println(strings.Join(str, ""))
	return strings.Join(str, "")
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	sum := sumOfLetters(n)
	fmt.Println(sum)
}
