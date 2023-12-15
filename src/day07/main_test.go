package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetHandValue(t *testing.T) {
	tcs := []struct {
		desc         string
		input        string
		getHandValue func(string) int
		expected     int
	}{
		{
			desc:         "expected Five of a kind",
			getHandValue: getHandValuePart1,
			input:        "AAAAA",
			expected:     7,
		},
		{
			desc:         "expected Four of a kind",
			getHandValue: getHandValuePart1,
			input:        "AA8AA",
			expected:     6,
		},
		{
			desc:         "expected Full house",
			getHandValue: getHandValuePart1,
			input:        "23332",
			expected:     5,
		},
		{
			desc:         "expected Three of a kind",
			getHandValue: getHandValuePart1,
			input:        "TTT98",
			expected:     4,
		},
		{
			desc:         "expected Two pair",
			getHandValue: getHandValuePart1,
			input:        "23432",
			expected:     3,
		},
		{
			desc:         "expected One pair",
			getHandValue: getHandValuePart1,
			input:        "A23A4",
			expected:     2,
		},
		{
			desc:         "expected High card",
			getHandValue: getHandValuePart1,
			input:        "23456",
			expected:     1,
		},
		{
			desc:         "expected expected Four of a kind",
			getHandValue: getHandValuePart2,
			input:        "QJJQ2",
			expected:     6,
		},
		{
			desc:         "expected expected Four of a kind",
			getHandValue: getHandValuePart2,
			input:        "JJJJJ",
			expected:     7,
		},
		{
			desc:         "expected expected Four of a kind",
			getHandValue: getHandValuePart2,
			input:        "AAAAJ",
			expected:     7,
		},
		{
			desc:         "expected expected Four of a kind",
			getHandValue: getHandValuePart2,
			input:        "AAAJJ",
			expected:     7,
		},
		{
			desc:         "expected expected Four of a kind",
			getHandValue: getHandValuePart2,
			input:        "AJJJJ",
			expected:     7,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value := tc.getHandValue(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetRankedHands(t *testing.T) {
	tcs := []struct {
		desc          string
		inputHands    []hand
		expectedHands []hand
		getCardValue,
		getHandValue func(s string) (v int)
	}{
		{
			desc:         "expected sorted part 1",
			getCardValue: getCardValuePart1,
			getHandValue: getHandValuePart1,
			inputHands: []hand{
				{
					cards: "32T3K",
					bid:   765,
				},
				{
					cards: "T55J5",
					bid:   684,
				},
				{
					cards: "KK677",
					bid:   28,
				},
				{
					cards: "KTJJT",
					bid:   220,
				},
				{
					cards: "QQQJA",
					bid:   483,
				},
			},
			expectedHands: []hand{
				{
					cards: "32T3K",
					bid:   765,
				},
				{
					cards: "KTJJT",
					bid:   220,
				},
				{
					cards: "KK677",
					bid:   28,
				},
				{
					cards: "T55J5",
					bid:   684,
				},
				{
					cards: "QQQJA",
					bid:   483,
				},
			},
		},
		{
			desc:         "expected sorted part 2",
			getCardValue: getCardValuePart2,
			getHandValue: getHandValuePart2,
			inputHands: []hand{
				{
					cards: "32T3K",
					bid:   765,
				},
				{
					cards: "T55J5",
					bid:   684,
				},
				{
					cards: "KK677",
					bid:   28,
				},
				{
					cards: "KTJJT",
					bid:   220,
				},
				{
					cards: "QQQJA",
					bid:   483,
				},
			},
			expectedHands: []hand{
				{
					cards: "32T3K",
					bid:   765,
				},
				{
					cards: "KK677",
					bid:   28,
				},
				{
					cards: "T55J5",
					bid:   684,
				},
				{
					cards: "QQQJA",
					bid:   483,
				},
				{
					cards: "KTJJT",
					bid:   220,
				},
			},
		},
		{
			desc:         "expected sorted part 2",
			getCardValue: getCardValuePart2,
			getHandValue: getHandValuePart2,
			inputHands: []hand{
				{
					cards: "QJJQ2",
					bid:   123,
				},
				{
					cards: "JKKK2",
					bid:   456,
				},
			},
			expectedHands: []hand{
				{
					cards: "JKKK2",
					bid:   456,
				},
				{
					cards: "QJJQ2",
					bid:   123,
				},
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			hands := getRankedHands(tc.inputHands, tc.getCardValue, tc.getHandValue)
			if diff := cmp.Diff(tc.expectedHands, hands, cmp.AllowUnexported(hand{})); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetTotalWinnings(t *testing.T) {
	tcs := []struct {
		desc  string
		input []string
		getCardValue,
		getHandValue func(s string) (v int)
		expected int
	}{
		{
			desc: "expected 6440 = ((765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5)",
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			getCardValue: getCardValuePart1,
			getHandValue: getHandValuePart1,
			expected:     6440,
		},
		{
			desc: "expected 5905 = ((765 * 1 + 28 * 2 + 684 * 3 + 483 * 4 + 220 * 5)",
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			getCardValue: getCardValuePart2,
			getHandValue: getHandValuePart2,
			expected:     5905,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getTotalWinnings(tc.input, tc.getCardValue, tc.getHandValue)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
