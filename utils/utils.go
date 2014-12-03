package utils

//Sieve uses a sieve of Arosthesis to fill a slice of primes up to n
func Sieve(amount int) []int {
	//fill a slice on 1..n
	n := make([]int, amount-1)
	for i := range n {
		n[i] = i + 2
	}
	//Starting number for the tests
	number := 2
	//stop at sqrt(highest prime we will get)
	for number*number <= len(n)+2 {
		for i := number + number; i < len(n)+2; i += number {
			n[i-2] = 0
		}
		number++
		for n[number-2] == 0 {
			number++
		}
	}
	//filter out the primes
	ret := *new([]int)
	for _, v := range n {
		if v != 0 {
			ret = append(ret, v)
		}
	}
	return ret
}
