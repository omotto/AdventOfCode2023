package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetFirsLastDigits(t *testing.T) {
	tcs := []struct {
		desc        string
		inputString string
		numOnly     bool
		expectedVal int
		expectedErr bool
	}{
		{
			desc:        "no value found",
			inputString: "string",
			numOnly:     true,
			expectedVal: 0,
			expectedErr: false,
		},
		{
			desc:        "expected 12",
			inputString: "1abc2",
			numOnly:     true,
			expectedVal: 12,
			expectedErr: false,
		},
		{
			desc:        "expected 38",
			inputString: "pqr3stu8vwx",
			numOnly:     true,
			expectedVal: 38,
			expectedErr: false,
		},
		{
			desc:        "expected 15",
			inputString: "a1b2c3d4e5f",
			numOnly:     true,
			expectedVal: 15,
			expectedErr: false,
		},
		{
			desc:        "expected 77",
			inputString: "treb7uchet",
			numOnly:     true,
			expectedVal: 77,
			expectedErr: false,
		},
		//
		{
			desc:        "expected 29",
			inputString: "two1nine",
			numOnly:     false,
			expectedVal: 29,
			expectedErr: false,
		},
		{
			desc:        "expected 83",
			inputString: "eightwothree",
			numOnly:     false,
			expectedVal: 83,
			expectedErr: false,
		},
		{
			desc:        "expected 13",
			inputString: "abcone2threexyz",
			numOnly:     false,
			expectedVal: 13,
			expectedErr: false,
		},
		{
			desc:        "expected 24",
			inputString: "xtwone3four",
			numOnly:     false,
			expectedVal: 24,
			expectedErr: false,
		},
		{
			desc:        "expected 42",
			inputString: "4nineeightseven2",
			numOnly:     false,
			expectedVal: 42,
			expectedErr: false,
		},
		{
			desc:        "expected 14",
			inputString: "zoneight234",
			numOnly:     false,
			expectedVal: 14,
			expectedErr: false,
		},
		{
			desc:        "expected 76",
			inputString: "7pqrstsixteen",
			numOnly:     false,
			expectedVal: 76,
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := getFirsLastDigits(tc.inputString, tc.numOnly)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
					t.Errorf("digits has diff %s", diff)
				}
			}
		})
	}
}

func TestGetSumCalibrationNumbers(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		numOnly     bool
		expectedVal int
		expectedErr bool
	}{
		{
			desc:        "no value found in list",
			inputList:   []string{},
			numOnly:     true,
			expectedVal: 0,
			expectedErr: false,
		},
		{
			desc:        "no digits found in list",
			inputList:   []string{"gdgsfggsdg"},
			numOnly:     true,
			expectedVal: 0,
			expectedErr: false,
		},
		{
			desc:        "expected 142",
			inputList:   []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			numOnly:     true,
			expectedVal: 142,
			expectedErr: false,
		},
		{
			desc:        "expected 281",
			inputList:   []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
			numOnly:     false,
			expectedVal: 281,
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := getSumCalibrationNumbers(tc.inputList, tc.numOnly)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
					t.Errorf("sum has diff %s", diff)
				}
			}
		})
	}
}
