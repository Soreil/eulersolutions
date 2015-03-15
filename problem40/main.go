package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := concat(1000000)
	sum := (str[0] - '0') * (str[9] - '0') * (str[99] - '0') * (str[999] - '0') * (str[9999] - '0') * (str[99999] - '0') * (str[999999] - '0')
	fmt.Println(sum)
}

func concat(n int) string {
	var buffer bytes.Buffer
	for i := 1; i <= n; i++ {
		buffer.WriteString(fmt.Sprintf("%d", i))
	}
	return buffer.String()
}
