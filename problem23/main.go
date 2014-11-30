package main

import "fmt"

type numberType int

func (n numberType) String() string {
	switch n {
	case deficient:
		return "deficient"
	case perfect:
		return "perfect"
	case abundant:
		return "abundant"
	default:
		return "error"
	}
}

const (
	deficient numberType = iota
	perfect
	abundant
)

//used for both the sums of all numbers and individual numbers
type typedNumber struct {
	num int
	t   numberType
}

func (t typedNumber) String() string {
	return fmt.Sprintf("%d %s", t.num, t.t)
}

//all primes lower than sqrt of lim
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
	43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107,
	109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167}

//largest number that can't be sum of two abundant numbers
var lim = 28123

func main() {
	nums := make([]typedNumber, lim) //all numbers and their type
	abundantNums := *new([]int)      //all of the abundant numbers
	for i := range nums {
		nums[i].num = i
		nums[i].t = sumType(divisorSum(nums[i].num), nums[i].num)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i].t == abundant {
			abundantNums = append(abundantNums, nums[i].num)
		}
	}
	twoAbundants := *new([]int)
	noTwoAbundants := *new([]int)
	//check all numbers that can be the sum of two abundant numbers
	for i := 1; i <= 28123; i++ {
		//label to exit in case of success of inner loop
	numberFound:
		//Check all numbers lower than the number we want to know the state of
		for j := 0; abundantNums[j] < i; j++ {
			for k := 0; abundantNums[j]+abundantNums[k] <= i; k++ {
				if abundantNums[j]+abundantNums[k] == i {
					twoAbundants = append(twoAbundants, i)
					break numberFound
				}
			}
		}
		//if Not possible to form i from two abundant sumbers
		if len(twoAbundants) > 0 {
			if twoAbundants[len(twoAbundants)-1] != i {
				noTwoAbundants = append(noTwoAbundants, i)
			}
		} else {
			noTwoAbundants = append(noTwoAbundants, i)
		}
	}
	fmt.Println("Amount of numbers that can't be formed from two abundant numbers:", len(noTwoAbundants))

	sum := 0
	for _, v := range noTwoAbundants {
		sum += v
	}

	fmt.Println("Sum of all numbers that can't be formed from two abundant numbers:", sum)
}

func divisorSum(n int) int {
	sum := 1
	x := n
	for i := range primes {
		if primes[i] > x/primes[i] { //sqrt reached,return
			break
		}
		if x%primes[i] == 0 {
			sub := primes[i] + 1 //increase power
			x /= primes[i]       //lower value to check against
			for x%primes[i] == 0 {
				x /= primes[i]          //lower value to check against
				sub = sub*primes[i] + 1 //Keep increasing power
			}
			sum *= sub //add to the divisor sum
		}
	}
	if x > 1 { // x is prime because it couldn't be lowered to 1
		sum *= x + 1
	}
	return sum - n //the number itself does not count as a proper divisor

}

func sumType(sum, number int) numberType {
	if sum > number {
		return abundant
	} else if sum == number {
		return perfect
	} else {
		return deficient
	}
}
