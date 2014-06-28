package main

import "fmt"
import "os"
import "strconv"

func main() {
	n, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
