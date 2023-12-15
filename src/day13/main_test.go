package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetMirrorNodes(t *testing.T) {
	tcs := []struct {
		desc           string
		input          []string
		expectedH      int
		expectedV      int
		getMirrorNodes func(p []string) (v, h int)
	}{
		{
			desc: "expected h:4, v:0 ",
			input: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			getMirrorNodes: getMirrorNodes,
			expectedH:      4,
			expectedV:      0,
		},
		{
			desc: "expected h:0, v:5 ",
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			getMirrorNodes: getMirrorNodes,
			expectedH:      0,
			expectedV:      5,
		},
		{
			desc: "Part2: expected h:3, v:0 ",
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			getMirrorNodes: getMirrorNodes2,
			expectedH:      3,
			expectedV:      0,
		},
		{
			desc: "Part2: expected h:1, v:0 ",
			input: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			getMirrorNodes: getMirrorNodes2,
			expectedH:      1,
			expectedV:      0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, h := tc.getMirrorNodes(tc.input)
			if diff := cmp.Diff(tc.expectedV, v); diff != "" {
				t.Errorf("vertical value has diff %s", diff)
			}
			if diff := cmp.Diff(tc.expectedH, h); diff != "" {
				t.Errorf("horizontal value has diff %s", diff)
			}
		})
	}
}
