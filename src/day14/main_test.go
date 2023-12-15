package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetTotalLoad(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 136 ",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: 136,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getTotalLoad(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("vertical value has diff %s", diff)
			}
		})
	}
}

func TestGetTotalLoad3(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		cycles   int
		expected int
	}{
		{
			desc: "expected 64 ",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			cycles:   1000000000,
			expected: 64,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getTotalLoad3(tc.input, tc.cycles)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("vertical value has diff %s", diff)
			}
		})
	}
}
