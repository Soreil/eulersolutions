package main

import "fmt"

const (
	_ = iota
	mon
	tue
	wed
	thu
	fri
	sat
	sun
)

const (
	_ = iota
	jan
	feb
	mar
	apr
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}

func main() {
	sundays := 0
	for year, day := 1901, mon; year < 2000; year++ {
		if isLeapYear(year) {
			switch day {
			case mon:
				sundays += 2
			case tue:
				sundays += 1
			case wed:
				sundays += 2
			case thu:
				sundays += 2
			case fri:
				sundays += 2
			case sat:
				sundays += 1
			case sun:
				sundays += 2
			}
			day += 2
		} else {
			switch day {
			case mon:
				sundays += 2
			case tue:
				sundays += 2
			case wed:
				sundays += 1
			case thu:
				sundays += 3
			case fri:
				sundays += 1
			case sat:
				sundays += 2
			case sun:
				sundays += 1
			}
			day += 1
		}
		day %= 7
		if day == 0 {
			day = sun
		}
	}
	fmt.Println("Sundays:", sundays)
}
