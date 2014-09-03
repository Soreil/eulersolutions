package main

import "fmt"

func biggest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tLen := 15
	triangle := make([][]int, tLen)
	triangle[0] = []int{75}
	triangle[1] = []int{95, 64}
	triangle[2] = []int{17, 47, 82}
	triangle[3] = []int{18, 35, 87, 10}
	triangle[4] = []int{20, 4, 82, 47, 65}
	triangle[5] = []int{19, 1, 23, 75, 3, 34}
	triangle[6] = []int{88, 2, 77, 73, 7, 63, 67}
	triangle[7] = []int{99, 65, 4, 28, 6, 16, 70, 92}
	triangle[8] = []int{41, 41, 26, 56, 83, 40, 80, 70, 33}
	triangle[9] = []int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
	triangle[10] = []int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
	triangle[11] = []int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
	triangle[12] = []int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
	triangle[13] = []int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
	triangle[14] = []int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}

	for i := tLen - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			triangle[i-1][j] += biggest(triangle[i][j], triangle[i][j+1])
		}
	}
	fmt.Println(triangle[0][0])
}
