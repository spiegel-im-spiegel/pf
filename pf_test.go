package pf

import (
	"os"
	"testing"
	"time"
)

type pfTest struct { //test case for GetPremiumFriday method
	inputY  int //input: Year
	inputM  int //input: Month
	outputD int //output Day of premium friday
}

var (
	pfTests []pfTest //test cases for GetPremiumFriday method
)

func TestMain(m *testing.M) {
	pfTests = []pfTest{
		{2016, 1, 0},
		{2017, 0, 0},
		{2017, 1, 0},
		{2017, 2, 24},
		{2017, 3, 31},
		{2017, 12, 29},
		{2017, 13, 0},
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestGetPremiumFriday(t *testing.T) {
	for _, testCase := range pfTests {
		result, _ := GetPremiumFriday(testCase.inputY, (time.Month)(testCase.inputM))
		if result != testCase.outputD {
			t.Errorf("Day of premium friday of (%v, %v) = %d, want %d.", testCase.inputY, testCase.inputM, result, testCase.outputD)
		}
	}
}
