package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParseLine2(t *testing.T) {
	tcs := []struct {
		desc              string
		input             string
		expectedDirection uint8
		expectedNumTiles  int
	}{
		{
			desc:              "expected #70c710 = R 461937",
			input:             "R 6 (#70c710)",
			expectedDirection: 'R',
			expectedNumTiles:  461937,
		},
		{
			desc:              "expected #0dc571 = D 56407",
			input:             "D 5 (#0dc571)",
			expectedDirection: 'D',
			expectedNumTiles:  56407,
		},
		{
			desc:              "expected #5713f0 = R 356671",
			input:             "L 2 (#5713f0)",
			expectedDirection: 'R',
			expectedNumTiles:  356671,
		},
		{
			desc:              "expected #d2c081 = D 863240",
			input:             "D 2 (#d2c081)",
			expectedDirection: 'D',
			expectedNumTiles:  863240,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			d, n, _ := parseLine2(tc.input)
			if diff := cmp.Diff(tc.expectedDirection, d); diff != "" {
				t.Errorf("direction has diff %s", diff)
			}
			if diff := cmp.Diff(tc.expectedNumTiles, n); diff != "" {
				t.Errorf("num tiles has diff %s", diff)
			}
		})
	}
}

func TestGetNumFilledTiles(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 62",
			input: []string{
				"R 6 (#70c710)",
				"D 5 (#0dc571)",
				"L 2 (#5713f0)",
				"D 2 (#d2c081)",
				"R 2 (#59c680)",
				"D 2 (#411b91)",
				"L 5 (#8ceee2)",
				"U 2 (#caa173)",
				"L 1 (#1b58a2)",
				"U 2 (#caa171)",
				"R 2 (#7807d2)",
				"U 3 (#a77fa3)",
				"L 2 (#015232)",
				"U 2 (#7a21e3)",
			},
			expected: 62,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getNumFilledTiles(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetNumFilledTiles2(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 952408144115",
			input: []string{
				"R 6 (#70c710)",
				"D 5 (#0dc571)",
				"L 2 (#5713f0)",
				"D 2 (#d2c081)",
				"R 2 (#59c680)",
				"D 2 (#411b91)",
				"L 5 (#8ceee2)",
				"U 2 (#caa173)",
				"L 1 (#1b58a2)",
				"U 2 (#caa171)",
				"R 2 (#7807d2)",
				"U 3 (#a77fa3)",
				"L 2 (#015232)",
				"U 2 (#7a21e3)",
			},
			expected: 952408144115,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getNumFilledTiles2(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
