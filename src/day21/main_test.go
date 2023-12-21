package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetGardenPlotsFilled(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		steps    int
		expected int
	}{
		{
			desc: "expected 2 for 1 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    1,
			expected: 2,
		},
		{
			desc: "expected 4 for 2 steps",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    2,
			expected: 4,
		},
		{
			desc: "expected 6 for 3 steps",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    3,
			expected: 6,
		},
		{
			desc: "expected 16",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    6,
			expected: 16,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getGardenPlotsFilled(tc.input, tc.steps)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetGardenPlotsFilled2(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		steps    int
		expected int
	}{
		{
			desc: "expected 16 for 6 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    6,
			expected: 16,
		},
		{
			desc: "expected 50 for 10 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    10,
			expected: 50,
		},
		{
			desc: "expected 1594 for 50 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    50,
			expected: 1594,
		},
		{
			desc: "expected 6536 for 100 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    100,
			expected: 6536,
		},
		{
			desc: "expected 167004 for 500 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    500,
			expected: 167004,
		},
		{
			desc: "expected 668697 for 1000 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    1000,
			expected: 668697,
		},
		{
			desc: "expected 16733044 for 5000 step",
			input: []string{
				"...........",
				".....###.#.",
				".###.##..#.",
				"..#.#...#..",
				"....#.#....",
				".##..S####.",
				".##..#...#.",
				".......##..",
				".##.#.####.",
				".##..##.##.",
				"...........",
			},
			steps:    5000,
			expected: 16733044,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getGardenPlotsFilled2(tc.input, tc.steps)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
