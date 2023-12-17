package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFindShortestPathTemperature(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 102",
			input: []string{
				"2413432311323",
				"3215453535623",
				"3255245654254",
				"3446585845452",
				"4546657867536",
				"1438598798454",
				"4457876987766",
				"3637877979653",
				"4654967986887",
				"4564679986453",
				"1224686865563",
				"2546548887735",
				"4322674655533",
			},
			expected: 102,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := findShortestPathTemperature(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestFindShortestPath(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 102",
			input: []string{
				"2413432311323",
				"3215453535623",
				"3255245654254",
				"3446585845452",
				"4546657867536",
				"1438598798454",
				"4457876987766",
				"3637877979653",
				"4654967986887",
				"4564679986453",
				"1224686865563",
				"2546548887735",
				"4322674655533",
			},
			expected: 102,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := findShortestPath(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestFindShortestPath2(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		min, max int
		expected int
	}{
		{
			desc: "expected 102",
			input: []string{
				"2413432311323",
				"3215453535623",
				"3255245654254",
				"3446585845452",
				"4546657867536",
				"1438598798454",
				"4457876987766",
				"3637877979653",
				"4654967986887",
				"4564679986453",
				"1224686865563",
				"2546548887735",
				"4322674655533",
			},
			min:      0,
			max:      3,
			expected: 102,
		},
		{
			desc: "expected 71",
			input: []string{
				"111111111111",
				"999999999991",
				"999999999991",
				"999999999991",
				"999999999991",
			},
			min:      4,
			max:      10,
			expected: 71,
		},
		{
			desc: "expected 94",
			input: []string{
				"2413432311323",
				"3215453535623",
				"3255245654254",
				"3446585845452",
				"4546657867536",
				"1438598798454",
				"4457876987766",
				"3637877979653",
				"4654967986887",
				"4564679986453",
				"1224686865563",
				"2546548887735",
				"4322674655533",
			},
			min:      4,
			max:      10,
			expected: 94,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := findShortestPath2(tc.input, tc.min, tc.max)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
