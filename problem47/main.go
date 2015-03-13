package main

/*Factor every number until a 4x4 is found.
Powers do NOT count, run a unique over the result slice
*/

import "fmt"
import "os"
import "strconv"

//TODO:(sjon) make factor utilize a precomputed prime list
func factor(n int) []int {
	s := make([]int, 0)
	d := 2
	for n > 1 {
		for n%d == 0 {
			s = append(s, d)
			n /= d
		}
		d++
	}
	return s
}

//Removes duplicates from an integer slice
func uniq(s []int) (ret []int) {
	vOld := -1
	for _, v := range s {
		if v != vOld {
			ret = append(ret, v)
			vOld = v
		}
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	lim, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		panic(err)
	}
	target := 0
	for i := 2; i < int(lim); i++ {
		s := uniq(factor(i))
		if len(s) == 4 {
			target++
		} else {
			target = 0
		}
		if target == 4 {
			fmt.Println(i)
			break
		}
	}
}
