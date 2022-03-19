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

func compareMaps() {
	for d1 := range WithoutLibMap {
		if _, ok := UsingLibMap[d1]; !ok {
			fmt.Println(d1, " only in withoutlib")
		}
	}
	for d2 := range UsingLibMap {
		if _, ok := WithoutLibMap[d2]; !ok {
			fmt.Println(d2, " only in UsingLibMap")
		}
	}
}
