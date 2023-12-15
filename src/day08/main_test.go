package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSteps(t *testing.T) {
	tcs := []struct {
		desc  string
		input []string
		getCardValue,
		getHandValue func(s string) (v int)
		expected int
	}{
		{
			desc: "expected 2 steps = AAA -R- CCC -L- ZZZ",
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 2,
		},
		{
			desc: "expected 6 steps = AAA -L- BBB -L- AAA -R- BBB -L- AAA -L- BBB -R- ZZZ",
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getSteps(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}

func TestGetSteps2(t *testing.T) {
	tcs := []struct {
		desc  string
		input []string
		getCardValue,
		getHandValue func(s string) (v int)
		expected int
	}{
		{
			desc: "expected 6 steps",
			input: []string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			expected: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			value, _ := getSteps2(tc.input)
			if diff := cmp.Diff(tc.expected, value); diff != "" {
				t.Errorf("values has diff %s", diff)
			}
		})
	}
}
