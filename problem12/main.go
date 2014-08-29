package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var primes = [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241}

func factors(n int64, c chan int64) {
	count := int64(0)
	for i := range primes {
		if n%primes[i] == 0 && primes[i] < int64(math.Sqrt(float64(n))) {
			for j := int64(1); j <= 500/primes[i]; j++ {
				if n%(primes[i]*j) == 0 {
					count++
				}
			}
		}
	}
	c <- count
}

func main() {
	tmp, err := strconv.Atoi(os.Args[1])
	n := int64(tmp)
	if err != nil {
		panic(err)
	}
	trianglenumber := int64(0)
	c := make(chan int64)
	done := make(chan bool)
	go func() {
		tmp := int64(0)
		count := int64(0)
		for count < n {
			if tmp = <-c; tmp > count {
				count = tmp
				fmt.Println(count)
				if count > 500 {
					done <- true
				}
			}
		}
	}()
	for i := int64(1); i <= n; i++ {
		trianglenumber += i
		go factors(trianglenumber, c)
	}
	<-done
}
