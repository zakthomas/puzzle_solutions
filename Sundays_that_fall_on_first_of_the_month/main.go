package main

import (
	"fmt"
	"os"
	"time"
)

// 1 Jan 1900 was a Monday.
// Thirty days has September,April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine. And on leap years, twenty-nine.

// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

var (
	daysOfTheWeek = []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	UsingLibMap   = make(map[string]string)
	WithoutLibMap = make(map[string]string)
)

type DayINFO struct {
	year    int
	month   time.Month
	date    int
	weekDay time.Weekday
}

// 1 Jan 1900 was a Monday.
var REFERENCE = DayINFO{year: 1900, month: 01, date: 01, weekDay: 1}

func main() {

	Init()

	startYear := 1901
	endYear := 2000

	fmt.Printf("INPUT : %d-%d\n", startYear, endYear)
	counter1 := UsingLib(startYear, endYear, time.Sunday)
	fmt.Printf("UsingLib   : Sundays that fell on the first of the month between %d and %d is %d\n", startYear, endYear, counter1)

	counter2 := WithoutLib(startYear, endYear, time.Sunday)
	fmt.Printf("withoutLib : Sundays that fell on the first of the month between %d and %d is %d\n", startYear, endYear, counter2)

}

func Init() {
	//Verify Reference
	referenceDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%02d", REFERENCE.year, REFERENCE.month, REFERENCE.date))
	if err != nil {
		panic(err)
	}
	if referenceDate.Weekday() != REFERENCE.weekDay {
		panic("REFERENCE DATE IS WRONG!!!")
	}
}

func UsingLib(StartYear, EndYear int, checkDay time.Weekday) int {
	startDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", StartYear))
	if err != nil {
		panic(err)
	}

	endDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-12-31", EndYear))
	if err != nil {
		panic(err)
	}

	counter := 0
	currentDate := startDate
	for !currentDate.After(endDate) {
		if currentDate.Day() == 1 && currentDate.Weekday() == checkDay {
			//fmt.Println("L:", currentDate.Format("02-01-2006"), " is a Sunday")
			UsingLibMap[currentDate.Format("02/01/2006")] = "sunday"
			counter++
		}
		currentDate = currentDate.AddDate(0, 1, 0)
	}
	return counter
}

func WithoutLib(StartYear, EndYear int, checkDay time.Weekday) int {

	counter := 0
	//Expected outcome is monthStart = jan 1 of start year
	monthStart := calculateStartDate(StartYear)

	for year := StartYear; year <= EndYear; year++ {
		monthList, _ := loadMonths(year)
		for m, val := range monthList {
			//if monthStart%7 == 0 { //This would work only for sundays because 0 +1 +2 +3 +4 +5 +6 +7 = 0 -1 -2 -3 -4 -5 -6 -7
			if findDay(monthStart%7) == checkDay {
				counter++
				//fmt.Printf("W: 01/%02d/%d is a Sunday\n", m+1, year)
				WithoutLibMap[fmt.Sprintf("01/%02d/%d", m+1, year)] = checkDay.String()
			}
			if year == EndYear && m == 11 { //Adding 31 december days will result in jan of next year
				continue
			} else {
				monthStart += val
			}
		}
	}
	return counter
}

func loadMonths(year int) ([]int, int) {
	leapYear := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	regularYear := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	leapYearTotal := 366
	regularYearTotal := 365

	if year%100 == 0 {
		if year%400 == 0 {
			//fmt.Println("\nLeap year: ", year)
			return leapYear, leapYearTotal
		}
	} else if year%4 == 0 {
		//fmt.Println("\nLeap year: ", year)
		return leapYear, leapYearTotal
	}
	return regularYear, regularYearTotal
}

func findDay(i int) time.Weekday {
	if i < 0 {
		i = 7 + i
	}
	return daysOfTheWeek[i]
}

func calculateStartDate(start int) int {

	monthStartDay := int(REFERENCE.weekDay) + 1 - REFERENCE.date
	switch {
	case start < REFERENCE.year:
		daysBack := 0
		for preload := REFERENCE.year - 1; preload >= start; preload-- {
			monthList, _ := loadMonths(preload)
			for i := 11; i >= 0; i-- {
				monthStartDay -= monthList[i]
				//fmt.Printf("\n01/%02d/%d > Calculted: %4d >> DaysToPrevious1st:%d TotalDaysReduced:%5d DayOfWeekExpected:%10s:%10s = %t", i+1, preload, monthStartDay, monthList[i], daysBack, dayOfTheWeek(preload, i+1, 1), findDay(monthStartDay%7), dayOfTheWeek(preload, i+1, 1) == findDay(monthStartDay%7))
				daysBack += monthList[i]
			}
			//fmt.Println("")
		}
	case start > REFERENCE.year:
		for preload := REFERENCE.year; preload < start; preload++ {
			monthList, _ := loadMonths(preload)
			for _, val := range monthList {
				monthStartDay += val
			}
		}
	}
	if findDay(monthStartDay%7) != dayOfTheWeek(start, 1, 1) {
		fmt.Printf("\n\nERROR  : Jan 1 of %d = %s (%d) want %s\n\n", start, findDay(monthStartDay%7), monthStartDay%7, dayOfTheWeek(start, 1, 1))
		os.Exit(1)
	}

	//fmt.Printf("\nSTARTING TO COUNT FROM ------------Jan 1 of %d = %s (%d)\n", start, findDay(monthStartDay%7), monthStartDay%7)
	return monthStartDay
}
