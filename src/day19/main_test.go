package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetSumRatingNumbers(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 19114",
			input: []string{
				"px{a<2006:qkq,m>2090:A,rfg}",
				"pv{a>1716:R,A}",
				"lnx{m>1548:A,A}",
				"rfg{s<537:gd,x>2440:R,A}",
				"qs{s>3448:A,lnx}",
				"qkq{x<1416:A,crn}",
				"crn{x>2662:A,R}",
				"in{s<1351:px,qqz}",
				"qqz{s>2770:qs,m<1801:hdj,R}",
				"gd{a>3333:R,R}",
				"hdj{m>838:A,pv}",
				"",
				"{x=787,m=2655,a=1222,s=2876}",
				"{x=1679,m=44,a=2067,s=496}",
				"{x=2036,m=264,a=79,s=2244}",
				"{x=2461,m=1339,a=466,s=291}",
				"{x=2127,m=1623,a=2188,s=1013}",
			},
			expected: 19114,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getSumRatingNumbers(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}

func TestGetAllCombinations(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc: "expected 167409079868000",
			input: []string{
				"px{a<2006:qkq,m>2090:A,rfg}",
				"pv{a>1716:R,A}",
				"lnx{m>1548:A,A}",
				"rfg{s<537:gd,x>2440:R,A}",
				"qs{s>3448:A,lnx}",
				"qkq{x<1416:A,crn}",
				"crn{x>2662:A,R}",
				"in{s<1351:px,qqz}",
				"qqz{s>2770:qs,m<1801:hdj,R}",
				"gd{a>3333:R,R}",
				"hdj{m>838:A,pv}",
				"",
				"{x=787,m=2655,a=1222,s=2876}",
				"{x=1679,m=44,a=2067,s=496}",
				"{x=2036,m=264,a=79,s=2244}",
				"{x=2461,m=1339,a=466,s=291}",
				"{x=2127,m=1623,a=2188,s=1013}",
			},
			expected: 167409079868000,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			v, _ := getAllCombinations(tc.input)
			if diff := cmp.Diff(tc.expected, v); diff != "" {
				t.Errorf("value has diff %s", diff)
			}
		})
	}
}
