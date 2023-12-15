package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput(t *testing.T) {
	tcs := []struct {
		desc        string
		inputString string
		expectedVal []cube
		expectedId  int
		expectedErr bool
	}{
		{
			desc:        "expected 1 [{3 4 0} {6 1 2} {0 0 2}] {blue, red, green}",
			inputString: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expectedVal: []cube{
				{
					blue:  3,
					red:   4,
					green: 0,
				},
				{
					blue:  6,
					red:   1,
					green: 2,
				},
				{
					blue:  0,
					red:   0,
					green: 2,
				},
			},
			expectedId:  1,
			expectedErr: false,
		},
		{
			desc:        "expected 2 [{1 0 2} {4 1 3} {1 0 1}] {blue, red, green}",
			inputString: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expectedVal: []cube{
				{
					blue:  1,
					red:   0,
					green: 2,
				},
				{
					blue:  4,
					red:   1,
					green: 3,
				},
				{
					blue:  1,
					red:   0,
					green: 1,
				},
			},
			expectedId:  2,
			expectedErr: false,
		},
		{
			desc:        "expected 3 [{6 20 8} {5 4 13} {0 1 5}] {blue, red, green}",
			inputString: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expectedVal: []cube{
				{
					blue:  6,
					red:   20,
					green: 8,
				},
				{
					blue:  5,
					red:   4,
					green: 13,
				},
				{
					blue:  0,
					red:   1,
					green: 5,
				},
			},
			expectedId:  3,
			expectedErr: false,
		},
		{
			desc:        "expected 4 [{6 3 1} {0 6 3} {15 14 3}] {blue, red, green}",
			inputString: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expectedVal: []cube{
				{
					blue:  6,
					red:   3,
					green: 1,
				},
				{
					blue:  0,
					red:   6,
					green: 3,
				},
				{
					blue:  15,
					red:   14,
					green: 3,
				},
			},
			expectedId:  4,
			expectedErr: false,
		},
		{
			desc:        "expected 5 [{1 6 3} {2 1 2}] {blue, red, green}",
			inputString: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expectedVal: []cube{
				{
					blue:  1,
					red:   6,
					green: 3,
				},
				{
					blue:  2,
					red:   1,
					green: 2,
				},
			},
			expectedId:  5,
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			id, val, err := parseInput(tc.inputString)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expectedId, id); diff != "" {
					t.Errorf("id has diff %s", diff)
				}
				if diff := cmp.Diff(tc.expectedVal, val, cmp.AllowUnexported(cube{})); diff != "" {
					t.Errorf("values has diff %s", diff)
				}
			}
		})
	}
}

func TestGetMinimumSetOfCubes(t *testing.T) {
	tcs := []struct {
		desc     string
		inputVal []cube
		expected cube
	}{
		{
			desc: "expected {6 4 2} {blue, red, green}",
			inputVal: []cube{
				{
					blue:  3,
					red:   4,
					green: 0,
				},
				{
					blue:  6,
					red:   1,
					green: 2,
				},
				{
					blue:  0,
					red:   0,
					green: 2,
				},
			},
			expected: cube{
				blue:  6,
				red:   4,
				green: 2,
			},
		},
		{
			desc: "expected {4 1 3} {blue, red, green}",
			inputVal: []cube{
				{
					blue:  1,
					red:   0,
					green: 2,
				},
				{
					blue:  4,
					red:   1,
					green: 3,
				},
				{
					blue:  1,
					red:   0,
					green: 1,
				},
			},
			expected: cube{
				blue:  4,
				red:   1,
				green: 3,
			},
		},
		{
			desc: "expected {6 20 13} {blue, red, green}",
			inputVal: []cube{
				{
					blue:  6,
					red:   20,
					green: 8,
				},
				{
					blue:  5,
					red:   4,
					green: 13,
				},
				{
					blue:  0,
					red:   1,
					green: 5,
				},
			},
			expected: cube{
				blue:  6,
				red:   20,
				green: 13,
			},
		},
		{
			desc: "expected {15 14 3} {blue, red, green}",
			inputVal: []cube{
				{
					blue:  6,
					red:   3,
					green: 1,
				},
				{
					blue:  0,
					red:   6,
					green: 3,
				},
				{
					blue:  15,
					red:   14,
					green: 3,
				},
			},
			expected: cube{
				blue:  15,
				red:   14,
				green: 3,
			},
		},
		{
			desc: "expected {2 6 3} {blue, red, green}",
			inputVal: []cube{
				{
					blue:  1,
					red:   6,
					green: 3,
				},
				{
					blue:  2,
					red:   1,
					green: 2,
				},
			},
			expected: cube{
				blue:  2,
				red:   6,
				green: 3,
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			val := getMinimumSetOfCubes(tc.inputVal)
			if diff := cmp.Diff(tc.expected, val, cmp.AllowUnexported(cube{})); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestCheckValidGame(t *testing.T) {
	tcs := []struct {
		desc      string
		inputList []cube
		valid     bool
	}{
		{
			desc: "[{3 4 0} {6 1 2} {0 0 2}] {blue, red, green} expected to be valid",
			inputList: []cube{
				{
					blue:  3,
					red:   4,
					green: 0,
				},
				{
					blue:  6,
					red:   1,
					green: 2,
				},
				{
					blue:  0,
					red:   0,
					green: 2,
				},
			},
			valid: true,
		},
		{
			desc: "[{1 0 2} {4 1 3} {1 0 1}] {blue, red, green} expected to be valid",
			inputList: []cube{
				{
					blue:  1,
					red:   0,
					green: 2,
				},
				{
					blue:  4,
					red:   1,
					green: 3,
				},
				{
					blue:  1,
					red:   0,
					green: 1,
				},
			},
			valid: true,
		},
		{
			desc: "[{6 20 8} {5 4 13} {0 1 5}] {blue, red, green} expected to be invalid",
			inputList: []cube{
				{
					blue:  6,
					red:   20,
					green: 8,
				},
				{
					blue:  5,
					red:   4,
					green: 13,
				},
				{
					blue:  0,
					red:   1,
					green: 5,
				},
			},
			valid: false,
		},
		{
			desc: "[{6 3 1} {0 6 3} {15 14 3}] {blue, red, green} expected to be invalid",
			inputList: []cube{
				{
					blue:  6,
					red:   3,
					green: 1,
				},
				{
					blue:  0,
					red:   6,
					green: 3,
				},
				{
					blue:  15,
					red:   14,
					green: 3,
				},
			},
			valid: false,
		},
		{
			desc: "[{1 6 3} {2 1 2}] {blue, red, green} expected to be valid",
			inputList: []cube{
				{
					blue:  1,
					red:   6,
					green: 3,
				},
				{
					blue:  2,
					red:   1,
					green: 2,
				},
			},
			valid: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := checkValidGame(tc.inputList, 13, 12, 14)
			if diff := cmp.Diff(tc.valid, got); diff != "" {
				t.Errorf("config has diff %s", diff)
			}
		})
	}
}

func TestGetSumValidIDGames(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
		expectedErr bool
	}{
		{
			desc: "expected 8",
			inputList: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedVal: 8,
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := getSumValidIDGames(tc.inputList)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
					t.Errorf("sum has diff %s", diff)
				}
			}
		})
	}
}

func TestGetSumPoweredCubes(t *testing.T) {
	tcs := []struct {
		desc        string
		inputList   []string
		expectedVal int
		expectedErr bool
	}{
		{
			desc: "expected 2286",
			inputList: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expectedVal: 2286,
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := getSumPoweredCubes(tc.inputList)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expectedVal, got); diff != "" {
					t.Errorf("sum has diff %s", diff)
				}
			}
		})
	}
}
