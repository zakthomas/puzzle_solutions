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
		input1   int
		input2   int
		output   int
		checkDay time.Weekday
	}{
		{2005, 2022, 0, -1},
		{1901, 2000, 171, time.Sunday}, //Sundays
		{1900, 2000, 0, -1},
		{1900, 1901, 0, -1},
		{1899, 1900, 0, -1},
		{1899, 1910, 0, -1},
		{1898, 1910, 0, -1},
		{1875, 1910, 0, -1},
		{1750, 1900, 0, -1},
		{1750, 1950, 0, -1},
		{1750, 1899, 0, -1},
	}

	inputCheckDays := []time.Weekday{time.Sunday, time.Wednesday, time.Friday}

	for _, checkDay := range inputCheckDays {
		for _, testCase := range tests {
			t.Run("Samples", func(t *testing.T) {
				//fmt.Printf("\nINPUT : %d-%d", testCase.input1, testCase.input2)
				resultA := UsingLib(testCase.input1, testCase.input2, checkDay)
				resultB := WithoutLib(testCase.input1, testCase.input2, checkDay)
				if resultA != resultB {
					compareMaps()
					t.Errorf("Between %d and %d : UsingLib =%d and WithoutLib =%d", testCase.input1, testCase.input2, resultA, resultB)
					fmt.Printf("Between %d and %d : UsingLib =%d and WithoutLib =%d\n\n", testCase.input1, testCase.input2, resultA, resultB)
					os.Exit(1)
				} else {
					if testCase.output != 0 && testCase.checkDay == checkDay {
						if testCase.output != resultA {
							t.Errorf("Between %d and %d : Expecting %d But got UsingLib =%d and WithoutLib =%d", testCase.input1, testCase.input2, testCase.output, resultA, resultB)
						} else {
							fmt.Printf("\nSUCCESS!! Between %d and %d : Expecting %d Got UsingLib =%d and WithoutLib =%d", testCase.input1, testCase.input2, testCase.output, resultA, resultB)
						}
					}
					//fmt.Printf(" >> %ss that fall on the first of the month = %5d", checkDay.String(), resultA)
				}
			})
		}
	}
	fmt.Println("\n----------------------------------------------------------------------------")
}
