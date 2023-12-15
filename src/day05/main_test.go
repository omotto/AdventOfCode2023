package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSeeds(t *testing.T) {
	tcs := []struct {
		desc       string
		inputSeeds string
		inputFunc  func(s string) (seeds []int, err error)
		expected   []int
	}{
		{
			desc:       "expected 79 14 55 13",
			inputSeeds: "seeds: 79 14 55 13",
			inputFunc:  getSeedsPart1,
			expected:   []int{79, 14, 55, 13},
		},
		{
			desc:       "expected 79, 80, 81, 82...92, 55, 56, 57...67",
			inputSeeds: "seeds: 79 14 55 13",
			inputFunc:  getSeedsPart2,
			expected:   []int{79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, _ := tc.inputFunc(tc.inputSeeds)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestCalculateLocation(t *testing.T) {
	_, m, _ := parseInput([]string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}, getSeedsPart1)

	tcs := []struct {
		desc        string
		inputSeed   int
		expectedLoc int
	}{
		{
			desc:        "expected 82",
			inputSeed:   79,
			expectedLoc: 82,
		},
		{
			desc:        "expected 43",
			inputSeed:   14,
			expectedLoc: 43,
		},
		{
			desc:        "expected 86",
			inputSeed:   55,
			expectedLoc: 86,
		},
		{
			desc:        "expected 35",
			inputSeed:   13,
			expectedLoc: 35,
		},
		{
			desc:        "expected 46",
			inputSeed:   82,
			expectedLoc: 46,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := calculateLocation(tc.inputSeed, m)
			if diff := cmp.Diff(tc.expectedLoc, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetLowerLocationNumber(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		seedFunc func(s string) (seeds []int, err error)
		expected int
	}{
		{
			desc:     "expected 35",
			seedFunc: getSeedsPart1,
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 35,
		},
		{
			desc:     "expected 46",
			seedFunc: getSeedsPart2,
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 46,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getLowerLocationNumber(tc.input, tc.seedFunc)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
