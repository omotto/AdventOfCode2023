package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSumLengths(t *testing.T) {
	tcs := []struct {
		desc        string
		input       []string
		inputFactor int
		expected    int
	}{
		{
			desc: "factor 2 -> expected 374",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			inputFactor: 2,
			expected:    374,
		},
		{
			desc: "factor 10 -> expected 1030",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			inputFactor: 10,
			expected:    1030,
		},
		{
			desc: "factor 100 -> expected 8410",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			inputFactor: 100,
			expected:    8410,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := getSumLengths(tc.input, tc.inputFactor)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
