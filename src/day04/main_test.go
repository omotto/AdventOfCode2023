package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCardStr(t *testing.T) {
	tcs := []struct {
		desc        string
		inputString string
		expected    Card
		expectedErr bool
	}{
		{
			desc:        "expected Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			inputString: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expected: Card{
				Id:      1,
				WinNums: []int{41, 48, 83, 86, 17},
				Nums:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expectedErr: false,
		},
		{
			desc:        "expected Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			inputString: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expected: Card{
				Id:      2,
				WinNums: []int{13, 32, 20, 16, 61},
				Nums:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expectedErr: false,
		},
		{
			desc:        "expected Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			inputString: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expected: Card{
				Id:      3,
				WinNums: []int{1, 21, 53, 59, 44},
				Nums:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			expectedErr: false,
		},
		{
			desc:        "expected Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			inputString: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			expected: Card{
				Id:      4,
				WinNums: []int{41, 92, 73, 84, 69},
				Nums:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			expectedErr: false,
		},
		{
			desc:        "expected Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			inputString: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			expected: Card{
				Id:      5,
				WinNums: []int{87, 83, 26, 28, 32},
				Nums:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expectedErr: false,
		},
		{
			desc:        "expected Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			inputString: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expected: Card{
				Id:      6,
				WinNums: []int{31, 18, 13, 56, 72},
				Nums:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			expectedErr: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			card, err := parseCardStr(tc.inputString)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error = %t, expectedErr %v", tc.expectedErr, err)
			}
			if !tc.expectedErr {
				if diff := cmp.Diff(tc.expected.Id, card.Id); diff != "" {
					t.Errorf("id has diff %s", diff)
				}
				if diff := cmp.Diff(tc.expected.Nums, card.Nums); diff != "" {
					t.Errorf("nums has diff %s", diff)
				}
				if diff := cmp.Diff(tc.expected.WinNums, card.WinNums); diff != "" {
					t.Errorf("winNums has diff %s", diff)
				}
			}
		})
	}
}

func TestGetNumWinnerNumbers(t *testing.T) {
	tcs := []struct {
		desc      string
		inputCard Card
		expected  []int
	}{
		{
			desc: "expected 48, 83, 17 and 86",
			inputCard: Card{
				Id:      1,
				WinNums: []int{41, 48, 83, 86, 17},
				Nums:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expected: []int{48, 83, 86, 17},
		},
		{
			desc: "expected 32 and 61",
			inputCard: Card{
				Id:      2,
				WinNums: []int{13, 32, 20, 16, 61},
				Nums:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expected: []int{32, 61},
		},
		{
			desc: "expected 1 and 21",
			inputCard: Card{
				Id:      3,
				WinNums: []int{1, 21, 53, 59, 44},
				Nums:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			expected: []int{1, 21},
		},
		{
			desc: "expected 84",
			inputCard: Card{
				Id:      4,
				WinNums: []int{41, 92, 73, 84, 69},
				Nums:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			expected: []int{84},
		},
		{
			desc: "expected no one",
			inputCard: Card{
				Id:      5,
				WinNums: []int{87, 83, 26, 28, 32},
				Nums:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expected: nil,
		},
		{
			desc: "expected no one",
			inputCard: Card{
				Id:      6,
				WinNums: []int{31, 18, 13, 56, 72},
				Nums:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			expected: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			values := getNumWinnerNumbers(tc.inputCard)
			if diff := cmp.Diff(tc.expected, values); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetPowFromCard(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "expected 8",
			input:    []int{48, 83, 86, 17},
			expected: 8,
		},
		{
			desc:     "expected 2",
			input:    []int{32, 61},
			expected: 2,
		},
		{
			desc:     "expected 2",
			input:    []int{1, 21},
			expected: 2,
		},
		{
			desc:     "expected 1",
			input:    []int{84},
			expected: 1,
		},
		{
			desc:     "expected 0",
			input:    nil,
			expected: 0,
		},
		{
			desc:     "expected 0",
			input:    nil,
			expected: 0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := getPowFromCard(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetSumOfAllCards(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 13",
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: 13,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getSumOfAllCards(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetNumWinnerCards(t *testing.T) {
	tcs := []struct {
		desc      string
		inputMap  []string
		inputLine int
		expected  int
	}{
		{
			desc: "expected 14 from line 0",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 0,
			expected:  14,
		},
		{
			desc: "expected 6 from line 1",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 1,
			expected:  6,
		},
		{
			desc: "expected 3 from line 2",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 2,
			expected:  3,
		},
		{
			desc: "expected 1 from line 3",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 3,
			expected:  1,
		},
		{
			desc: "expected 0 from line 4",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 4,
			expected:  0,
		},
		{
			desc: "expected 0 from line 5",
			inputMap: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			inputLine: 5,
			expected:  0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := getNumWinnerCards(tc.inputMap, tc.inputLine, 0)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetTotalWinnerCards(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 30",
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: 30,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := getTotalWinnerCards(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
