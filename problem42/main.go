package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	words, err := readFile("words.txt")
	if err != nil {
		panic(err)
	}
	var sum int
	for i := range words {
		if isTriangular(wordToInt(words[i])) {
			sum++
		}
	}
	fmt.Println(sum)
}

func isTriangular(n int) bool {
	sqrt := math.Sqrt(float64(8*n + 1))
	return math.Floor(sqrt) == sqrt
}

//Expected input is upper case ASCII
func wordToInt(s string) int {
	sum := 0
	for _, v := range s {
		sum += int(v) - 'A' + 1
	}
	return sum

}

//takes a filename and returns a two dimensional array of the ints
func readFile(fname string) ([]string, error) {
	//read file
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	//split words
	words := strings.Split(string(b), ",")

	for i := range words {
		words[i] = strings.TrimPrefix(words[i], "\"")
		words[i] = strings.TrimSuffix(words[i], "\"")
	}
	return words, nil
}
