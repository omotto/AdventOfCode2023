package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSumTilesEnergized(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 46",
			input: []string{
				`.|...\....`,
				`|.-.\.....`,
				`.....|-...`,
				`........|.`,
				`..........`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 46,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := getSumTilesEnergized(tc.input, 0, 0, Right)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
