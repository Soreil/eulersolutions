package main

func sieve(n []int64) []int64 {
	number := int64(2)
	for number*number <= int64(len(n)+2) {
		for i := number + number; i < int64(len(n)+2); i += number {
			n[i-2] = 0
		}
		number++
		for n[number-2] == 0 {
			number++
		}
	}
	return n[:len(n)-1]
}

func fill(n int64) []int64 {
	list := make([]int64, n)
	for i := int64(1); i < n; i++ {
		list[i-1] = i + 1
	}
	return list
}

func clean(list []int64) (cleaned []int64) {
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			cleaned = append(cleaned, list[i])
		}
	}
	return
}

func main() {
	list := fill(10000)
	list = sieve(list)
	fourDigits := list[1000:10000]
	fourDigits = clean(fourDigits)
}
