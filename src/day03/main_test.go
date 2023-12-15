package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetValidNumbers(t *testing.T) {
	tcs := []struct {
		desc     string
		inputMap []string
		expected []int
	}{
		{
			desc: "expected {467, 35, 633, 617, 592, 755, 664, 598}",
			inputMap: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: []int{467, 35, 633, 617, 592, 755, 664, 598},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			gnums := getValidNumbers(tc.inputMap)
			if diff := cmp.Diff(tc.expected, gnums); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumValidNumsFromMap(t *testing.T) {
	tcs := []struct {
		desc     string
		inputMap []string
		expected int
	}{
		{
			desc: "expected 4361",
			inputMap: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumValidNumsFromMap(tc.inputMap)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetGearNumbers(t *testing.T) {
	tcs := []struct {
		desc     string
		inputMap []string
		expected []int
	}{
		{
			desc: "expected {16345, 451490}",
			inputMap: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: []int{16345, 451490},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			gnums := getGearNumbers(tc.inputMap)
			if diff := cmp.Diff(tc.expected, gnums); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetGearValidNumsFromMap(t *testing.T) {
	tcs := []struct {
		desc     string
		inputMap []string
		expected int
	}{
		{
			desc: "expected 467835",
			inputMap: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 467835,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumGearNumsFromMap(tc.inputMap)
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
