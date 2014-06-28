package main

import "fmt"
import "os"
import "strconv"

func sumSqDif(n int) float64 {
	sqsum := float64(n*(n+1)) / 2
	sumsq := float64(n*(n+1)*(2*n+1)) / float64(6)
	return sqsum*sqsum - sumsq
}
func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(sumSqDif(int(n)))
}
