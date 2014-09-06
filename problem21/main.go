package main

import "fmt"

func divisors(n int64) (divisors []int64) {
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return
}

func sum(n []int64) (total int64) {
	for _, v := range n {
		total += v
	}
	return
}

func main() {
	var amiSum int64
	for i := 0; i < 10000; i++ {
		a := int64(i)
		tmp := sum(divisors(a))
		b := sum(divisors(tmp))
		if a == b && a != tmp {
			fmt.Println(a, tmp, " Are amicable numbers.")
			amiSum += a + tmp
		}
	}
	fmt.Println("sum of all Amicable numbers under n:", amiSum/2)
}
