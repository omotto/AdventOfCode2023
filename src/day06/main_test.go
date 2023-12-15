package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalculateDistance(t *testing.T) {
	tcs := []struct {
		desc         string
		inputTime    int
		inputMaxTime int
		expected     int
	}{
		{
			desc:         "for 0ms and 7ms max expected 0mm",
			inputTime:    0,
			inputMaxTime: 7,
			expected:     0,
		},
		{
			desc:         "for 1ms and 7ms max expected 6mm",
			inputTime:    1,
			inputMaxTime: 7,
			expected:     6,
		},
		{
			desc:         "for 2ms and 7ms max expected 10mm",
			inputTime:    2,
			inputMaxTime: 7,
			expected:     10,
		},
		{
			desc:         "for 3ms and 7ms max expected 12mm",
			inputTime:    3,
			inputMaxTime: 7,
			expected:     12,
		},
		{
			desc:         "for 4ms and 7ms max expected 12mm",
			inputTime:    4,
			inputMaxTime: 7,
			expected:     12,
		},
		{
			desc:         "for 5ms and 7ms max expected 10mm",
			inputTime:    5,
			inputMaxTime: 7,
			expected:     10,
		},
		{
			desc:         "for 6ms and 7ms max expected 6mm",
			inputTime:    6,
			inputMaxTime: 7,
			expected:     6,
		},
		{
			desc:         "for 7ms and 7ms max expected 0mm",
			inputTime:    7,
			inputMaxTime: 7,
			expected:     0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := calculateDistance(tc.inputTime, tc.inputMaxTime)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestCalculateWinners(t *testing.T) {
	tcs := []struct {
		desc          string
		inputDuration int
		inputMaxTime  int
		expected      int
	}{
		{
			desc:          "expected 4 winners",
			inputDuration: 9,
			inputMaxTime:  7,
			expected:      4,
		},
		{
			desc:          "expected 8 winners",
			inputDuration: 40,
			inputMaxTime:  15,
			expected:      8,
		},
		{
			desc:          "expected 9 winners",
			inputDuration: 200,
			inputMaxTime:  30,
			expected:      9,
		},
		{
			desc:          "expected 71503 winners",
			inputDuration: 940200,
			inputMaxTime:  71530,
			expected:      71503,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			values := calculateWinners(tc.inputMaxTime, tc.inputDuration)
			if diff := cmp.Diff(tc.expected, len(values)); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetResults(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 288 = (4 * 8 * 9)",
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 288,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getResults(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
