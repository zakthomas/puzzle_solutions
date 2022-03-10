package main

import (
	"fmt"
	"time"
)

// 1 Jan 1900 was a Monday.
// Thirty days has September,April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine. And on leap years, twenty-nine.

// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

func main() {

	startYear := 1901
	endYear := 2000

	usingLib(startYear, endYear)
	withoutLib(startYear, endYear)

}

func usingLib(startYear int, endYear int) {
	startDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", startYear))
	if err != nil {
		panic(err)
	}

	endDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-12-31", endYear))
	if err != nil {
		panic(err)
	}

	counter := 0
	currentDate := startDate
	for !currentDate.After(endDate) {
		if currentDate.Day() == 1 && currentDate.Weekday() == 0 {
			//fmt.Println(currentDate.Format("2006-01-02"), " is a Sunday")
			counter++
		}
		currentDate = currentDate.AddDate(0, 1, 0)
	}
	fmt.Printf("usingLib   : Sundays that fell on the first of the month between %s and %s is %d\n", startDate.Format("2006"), endDate.Format("2006"), counter)
}

func withoutLib(start int, end int) {
	counter := 0
	//Given reference = 1 monday (1900 jan 1)
	monthStart := 1 // Because monday%7=1
	for _, val := range loadMonths(1900) {
		monthStart += val
	}
	monthCounter := 0
	for year := start; year <= end; year++ {
		for _, val := range loadMonths(year) {
			monthCounter++
			monthStart += val
			if monthStart%7 == 0 {
				counter++
			}
		}
	}
	fmt.Printf("withoutLib : Sundays that fell on the first of the month between %d and %d is %d\n", start, end, counter)
}

func loadMonths(year int) []int {
	leapYear := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	regularYear := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%100 == 0 {
		if year%400 == 0 {
			return leapYear
		}
	} else if year%4 == 0 {
		return leapYear
	}
	return regularYear
}
