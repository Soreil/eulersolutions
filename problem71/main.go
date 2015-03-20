package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/Soreil/eulersolutions/utils"
)

/*
bounds: 2/5 - 3/7
Our goal is to approach 3/7
*/
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var primes []int

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	primes = utils.Sieve(10000)
	lower := 2 / 5.0
	upper := 3 / 7.0
	nLower := 2.0
	for d := 9.0; d < 1000000; d++ {
		for n := nLower + 1; n < d; n++ {
			if gcd(int(n), int(d), primes) == 1 {
				if n/d > lower && n/d < upper {
					lower = n / d
					nLower = n
					fmt.Println(lower, n, d)
				} else if n/d > upper {
					break
				}
			}
		}
	}
}

func gcd(a, b int, primes []int) int {
	gcd := 1
	for _, n := range primes {
		for i := 1; a%(n*i) == 0 && b%(n*i) == 0; i++ {
			gcd *= n
		}
		if a == 1 || b == 1 {
			break
		}
	}

	return gcd
}
