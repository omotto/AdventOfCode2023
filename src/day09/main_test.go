package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCalculatePrevNumber(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "expected -3 -> 0, 3, 6, 9, 12, 15",
			input:    []int{0, 3, 6, 9, 12, 15},
			expected: -3,
		},
		{
			desc:     "expected 0 -> 1, 3, 6, 10, 15, 21",
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: 0,
		},
		{
			desc:     "expected 5 -> 10, 13, 16, 21, 30, 45",
			input:    []int{10, 13, 16, 21, 30, 45},
			expected: 5,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := calculatePrevNumber(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestCalculateNextNumber(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "expected 18 -> 0, 3, 6, 9, 12, 15",
			input:    []int{0, 3, 6, 9, 12, 15},
			expected: 18,
		},
		{
			desc:     "expected 28 -> 1, 3, 6, 10, 15, 21",
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: 28,
		},
		{
			desc:     "expected 68 -> 10, 13, 16, 21, 30, 45",
			input:    []int{10, 13, 16, 21, 30, 45},
			expected: 68,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := calculateNextNumber(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetSumOASIS2(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 2 -> -3 + 0 + 5",
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 2,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getSumOASIS2(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
