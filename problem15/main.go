package main

import (
	"fmt"
	"os"
	"strconv"
)

func paths(grid [][]int64) (count int64) {
	yMax := len(grid)
	xMax := len(grid[0])
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if y == 0 {
				grid[y][x] += 1
			} else {
				grid[y][x] += grid[y-1][x]
			}
			if x == 0 {
				grid[y][x] += 1
			} else {
				grid[y][x] += grid[y][x-1]
			}

		}
	}
	for i := range grid {
		fmt.Println(grid[i])
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

//generate a y*x grid
func genGrid(y, x int) (grid [][]int64) {
	grid = make([][]int64, y)
	for i := range grid {
		grid[i] = make([]int64, x)
	}
	return
}

func main() {
	//provide the y and x(in this case we use a square so only y)
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(paths(genGrid(n, n)))
}
