package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetMultipliedSize(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 54",
			input: []string{
				"jqt: rhn xhk nvd",
				"rsh: frs pzl lsr",
				"xhk: hfx",
				"cmg: qnr nvd lhk bvb",
				"rhn: xhk bvb hfx",
				"bvb: xhk hfx",
				"pzl: lsr hfx nvd",
				"qnr: nvd",
				"ntq: jqt hfx bvb xhk",
				"nvd: lhk",
				"lsr: lhk",
				"rzs: qnr cmg lsr rsh",
				"frs: qnr lhk lsr",
			},
			expected: 54,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getMultipliedSize(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
