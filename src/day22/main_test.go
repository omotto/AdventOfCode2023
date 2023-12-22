package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSafelyDesintegratedBrisks(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 5",
			input: []string{
				"1,0,1~1,2,1",
				"0,0,2~2,0,2",
				"0,2,3~2,2,3",
				"0,0,4~0,2,4",
				"2,0,5~2,2,5",
				"0,1,6~2,1,6",
				"1,1,8~1,1,9",
			},
			expected: 5,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getSafelyDesintegratedBrisks(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetTotalFallen(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 7",
			input: []string{
				"1,0,1~1,2,1",
				"0,0,2~2,0,2",
				"0,2,3~2,2,3",
				"0,0,4~0,2,4",
				"2,0,5~2,2,5",
				"0,1,6~2,1,6",
				"1,1,8~1,1,9",
			},
			expected: 7,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getTotalFallen(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
