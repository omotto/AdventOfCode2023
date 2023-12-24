package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetCrossedHailstones(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		min, max float64
		expected int
	}{
		{
			desc: "expected 2",
			input: []string{
				"19, 13, 30 @ -2,  1, -2",
				"18, 19, 22 @ -1, -1, -2",
				"20, 25, 34 @ -2, -2, -4",
				"12, 31, 28 @ -1, -2, -1",
				"20, 19, 15 @  1, -5, -3",
			},
			min:      7.0,
			max:      27.0,
			expected: 2,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getCrossedHailstones(tc.input, tc.min, tc.max)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
