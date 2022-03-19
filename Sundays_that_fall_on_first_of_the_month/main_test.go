package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCounts(t *testing.T) {

	Init()

	var tests = []struct {
		input1     int
		input2     int
		fallOnDate int
		output     int
		checkDay   time.Weekday
	}{
		{2005, 2022, 25, 0, -1},
		{1901, 2000, 1, 171, time.Sunday}, //Sundays
		{1900, 2000, 25, 0, -1},
		{1900, 1901, 25, 0, -1},
		{1899, 1900, 25, 0, -1},
		{1899, 1910, 25, 0, -1},
		{1898, 1910, 25, 0, -1},
		{1875, 1910, 25, 0, -1},
		{1750, 1900, 25, 0, -1},
		{1750, 1950, 25, 0, -1},
		{1750, 1899, 25, 0, -1},
	}

	inputCheckDays := []time.Weekday{time.Sunday, time.Wednesday, time.Friday}

	for _, checkDay := range inputCheckDays {
		for _, testCase := range tests {
			t.Run("Samples", func(t *testing.T) {
				//fmt.Printf("\nINPUT : %d-%d", testCase.input1, testCase.input2)
				resultA := UsingLib(testCase.input1, testCase.input2, testCase.fallOnDate, checkDay)
				resultB := WithoutLib(testCase.input1, testCase.input2, testCase.fallOnDate, checkDay)
				if resultA != resultB {
					fmt.Printf("\nINPUT : %d-%d\n", testCase.input1, testCase.input2)
					compareMaps(checkDay)
					t.Errorf("\nBetween %d and %d : UsingLib =%d and WithoutLib =%d\n\n", testCase.input1, testCase.input2, resultA, resultB)
					fmt.Printf("\nBetween %d and %d : UsingLib =%d and WithoutLib =%d\n\n", testCase.input1, testCase.input2, resultA, resultB)
					os.Exit(1)
				} else {
					if testCase.output != 0 && testCase.checkDay == checkDay {
						if testCase.output != resultA {
							t.Errorf("Between %d and %d : Expecting %d But got UsingLib =%d and WithoutLib =%d", testCase.input1, testCase.input2, testCase.output, resultA, resultB)
						} else {
							fmt.Printf("\nSUCCESS!! Between %d and %d : Expecting %d Got UsingLib =%d and WithoutLib =%d", testCase.input1, testCase.input2, testCase.output, resultA, resultB)
						}
					}
					fmt.Printf("\n %d-%d > %ss that fall on the %d of the month = %5d", testCase.input1, testCase.input2, checkDay.String(), testCase.fallOnDate, resultA)
				}
			})
		}
		fmt.Println()
	}
	fmt.Println("\n----------------------------------------------------------------------------")
}
