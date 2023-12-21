package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMultipliedPulses(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 32000000",
			input: []string{
				"broadcaster -> a, b, c",
				"%a -> b",
				"%b -> c",
				"%c -> inv",
				"&inv -> a",
			},
			expected: 32000000,
		},
		{
			desc: "expected 11687500",
			input: []string{
				"broadcaster -> a",
				"%a -> inv, con",
				"&inv -> b",
				"%b -> con",
				"&con -> output",
			},
			expected: 11687500,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getMultipliedPulses(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
