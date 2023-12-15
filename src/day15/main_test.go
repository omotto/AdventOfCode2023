package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "expected 52 from HASH ",
			input:    "HASH",
			expected: 52,
		},
		{
			desc:     "expected 30 from rn=1 ",
			input:    "rn=1",
			expected: 30,
		},
		{
			desc:     "expected 253 from cm- ",
			input:    "cm-",
			expected: 253,
		},
		{
			desc:     "expected 97 from qp=3 ",
			input:    "qp=3",
			expected: 97,
		},
		{
			desc:     "expected 47 from cm=2 ",
			input:    "cm=2",
			expected: 47,
		},
		{
			desc:     "expected 14 from qp- ",
			input:    "qp-",
			expected: 14,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := calculateHash(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetSumHash(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "expected 1320 from HASH sum",
			input:    []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"},
			expected: 1320,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := getSumHash(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetFocusingPower(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "expected 145 from HASH sum",
			input:    []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"},
			expected: 145,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v := getFocusingPower(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
