package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(year int, month int, date int) time.Weekday {
	datestring := fmt.Sprintf("%d-%02d-%02d", year, month, date)
	givenDate, err := time.Parse("2006-01-02", datestring)
	if err != nil {
		panic(err)
	}
	return findDay(int(givenDate.Weekday()))
}

func compareMaps(checkDay time.Weekday) {
	for d1, value := range WithoutLibMap {
		if _, ok := UsingLibMap[d1]; !ok {
			if !(value.Weekday() == checkDay) {
				fmt.Printf("\nMismatch in WithoutLibMap :%s is not a %s", value.String(), checkDay.String())
			}
		}
	}
	for d2, value := range UsingLibMap {
		if _, ok := WithoutLibMap[d2]; !ok {
			if !(value.Weekday() == checkDay) {
				fmt.Printf("\nMismatch in UsingLibMap :%s is not a %s", value.String(), checkDay.String())
			}
		}
	}
}
