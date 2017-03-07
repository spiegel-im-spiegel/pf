package pf

import (
	"os"
	"testing"
)

type pfTest struct { //test case for GetPremiumFriday method
	inputY  int //input: Year
	inputM  int //input: Month
	outputD int //output Day of premium friday
}

type pfTest2 struct { //test case for IsPremiumFriday method
	inputY     int  //input: Year
	inputM     int  //input: Month
	inputD     int  //input Day
	outputBool bool //output: boolean
}

var (
	pfTests  []pfTest  //test cases for GetPremiumFriday method
	pfTests2 []pfTest2 //test cases for IsPremiumFriday method
)

func TestMain(m *testing.M) {
	pfTests = []pfTest{
		{2016, 1, 0},
		{2017, 0, 0}, //2016-12
		{2017, 1, 0},
		{2017, 2, 24},
		{2017, 3, 31},
		{2017, 12, 29},
		{2017, 13, 26}, //pf = 2018-01-26
	}

	pfTests2 = []pfTest2{
		{2017, 1, 27, false},
		{2017, 2, 17, false},
		{2017, 2, 23, false},
		{2017, 2, 24, true},
		{2017, 2, 25, false},
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestGetPremiumFriday(t *testing.T) {
	for _, testCase := range pfTests {
		dt := NewYearMonth(testCase.inputY, testCase.inputM)
		result, _ := dt.GetPremiumFriday()
		if result != testCase.outputD {
			t.Errorf("Day of premium friday of (%v, %v) = %d, want %d.", testCase.inputY, testCase.inputM, result, testCase.outputD)
		}
	}
	//Error case
	dt := NewDate(2017, 2, 24)
	result, _ := dt.GetPremiumFriday()
	if result != 0 {
		t.Errorf("Day of premium friday = %d, want %d.", result, 0)
	}
}

func TestIsPremiumFriday(t *testing.T) {
	for _, testCase := range pfTests2 {
		dt := NewDate(testCase.inputY, testCase.inputM, testCase.inputD)
		result := dt.IsPremiumFriday()
		if result != testCase.outputBool {
			t.Errorf("Day (%v-%v-%v) Premium Friday ? -> %v, want %v.", testCase.inputY, testCase.inputM, testCase.inputD, result, testCase.outputBool)
		}
	}
	//Error case
	dt := NewYearMonth(2017, 2)
	result := dt.IsPremiumFriday()
	if result != false {
		t.Errorf("Result = %v, want %v.", result, false)
	}
}
