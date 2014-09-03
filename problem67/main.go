package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func biggest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//takes a filename and returns a two dimensional array of the ints
func readFile(fname string) (nums [][]int, err error) {
	var tmp int //Couldn't be bothered to mess with the different scopes
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	nums = make([][]int, len(lines)-1)

	for i, _ := range lines {
		buf := strings.Split(lines[i], " ")
		for _, l := range buf {
			if i < 100 { //Split adds another line, we don't want this one.
				tmp, err = strconv.Atoi(l)
				if err != nil {
					panic(err)
				}
			} else {
				return
			}
			nums[i] = append(nums[i], tmp)
		}
	}
	return
}

func main() {
	tLen := 100
	// make the datastructure to save the triangle in
	triangle, err := readFile("p067_triangle.txt")
	if err != nil {
		panic(err)
	}

	//walk the triangle from bottom to top
	for i := tLen - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			triangle[i-1][j] += biggest(triangle[i][j], triangle[i][j+1])
		}
	}
	//print the top with the final value
	fmt.Println(triangle[0][0])
}
