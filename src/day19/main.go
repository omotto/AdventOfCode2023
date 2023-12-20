package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type part struct {
	x, m, a, s int
}

type wf map[string][]string

func parseInput(lines []string) (wf, []part, error) {
	var (
		parts       []part
		wfs         = make(wf)
		isWorkFlows = true
	)
	for _, line := range lines {
		if len(line) > 2 {
			if isWorkFlows {
				ws := strings.Split(line[:len(line)-1], "{")
				if len(ws) != 2 {
					return wfs, parts, errors.New("invalid input")
				}
				key := ws[0]
				ops := strings.Split(ws[1], ",")
				wfs[key] = ops
			} else {
				p := line[1 : len(line)-1]
				ps := strings.Split(p, ",")
				x := 0
				m := 0
				s := 0
				a := 0
				for i := 0; i < len(ps); i++ {
					pp := strings.Split(ps[i], "=")
					if len(pp) != 2 {
						return wfs, parts, errors.New("invalid input")
					}
					if value, err := strconv.Atoi(pp[1]); err != nil {
						return wfs, parts, err
					} else {
						switch pp[0] {
						case "x":
							x = value
						case "m":
							m = value
						case "a":
							a = value
						case "s":
							s = value
						default:
							return wfs, parts, errors.New("invalid input")
						}
					}
				}
				parts = append(parts, part{
					x: x,
					m: m,
					a: a,
					s: s,
				})
			}
		} else {
			isWorkFlows = !isWorkFlows
		}
	}
	return wfs, parts, nil
}

func evaluate(p part, ops []string) string {
	var (
		value     int
		param     string
		gotoKey   string
		condition uint8
	)
	for _, op := range ops {
		i := strings.Split(op, ":")
		if len(i) == 1 {
			return i[0]
		}
		gotoKey = i[1]
		j := strings.Split(i[0], "<")
		k := strings.Split(i[0], ">")
		if len(j) == 2 {
			value, _ = strconv.Atoi(j[1])
			param = j[0]
			condition = '<'
		} else {
			value, _ = strconv.Atoi(k[1])
			param = k[0]
			condition = '>'
		}
		switch param {
		case "x":
			if condition == '>' {
				if p.x > value {
					return gotoKey
				}
			} else {
				if p.x < value {
					return gotoKey
				}
			}
		case "m":
			if condition == '>' {
				if p.m > value {
					return gotoKey
				}
			} else {
				if p.m < value {
					return gotoKey
				}
			}
		case "a":
			if condition == '>' {
				if p.a > value {
					return gotoKey
				}
			} else {
				if p.a < value {
					return gotoKey
				}
			}
		case "s":
			if condition == '>' {
				if p.s > value {
					return gotoKey
				}
			} else {
				if p.s < value {
					return gotoKey
				}
			}
		}
	}
	panic("solution not found")
}

func getSumRatingNumbers(s []string) (int, error) {
	var total int = 0
	if wfs, parts, err := parseInput(s); err != nil {
		return 0, err
	} else {
		for _, p := range parts {
			wfKey := "in"
			exit := false
			for {
				r := evaluate(p, wfs[wfKey])
				switch r {
				case "A":
					total += p.x + p.m + p.s + p.a
					exit = true
				case "R":
					exit = true
				default:
					wfKey = r
				}
				if exit {
					break
				}
			}
		}
	}
	return total, nil
}

func parseWorkFlows(lines []string) (map[string]string, error) {
	var key string
	wfs := make(map[string]string)
	for _, line := range lines {
		if len(line) > 2 {
			ws := strings.Split(line[:len(line)-1], "{")
			if len(ws) != 2 {
				return wfs, errors.New("invalid input")
			}
			ops := strings.Split(ws[1], ",")
			if len(ops) == 2 {
				key := ws[0]
				wfs[key] = ws[1]
			} else {
				for i := 0; i < len(ops)-1; i++ {
					if i == 0 {
						key = ws[0]
					} else {
						key = fmt.Sprintf("%s%d", ws[0], i)
					}
					if i == len(ops)-2 {
						wfs[key] = fmt.Sprintf("%s,%s", ops[i], ops[i+1])
					} else {
						wfs[key] = fmt.Sprintf("%s,%s%d", ops[i], ws[0], i+1)
					}
				}
			}
		} else {
			break
		}
	}
	return wfs, nil
}

func getAccepted(wfKey string, xMin, xMax, mMin, mMax, aMin, aMax, sMin, sMax int, wfs map[string]string) int {
	if wfKey == "A" {
		return (xMax - xMin + 1) * (mMax - mMin + 1) * (aMax - aMin + 1) * (sMax - sMin + 1)
	} else if wfKey == "R" {
		return 0
	} else {
		acc := 0
		condition := wfs[wfKey]
		a := strings.Split(condition, ":")
		options := strings.Split(a[1], ",")
		trueOption := options[0]
		falseOption := options[1]
		param := a[0][0]
		symbol := a[0][1]
		value, _ := strconv.Atoi(a[0][2:])
		switch param {
		case 'x':
			if symbol == '>' {
				acc += getAccepted(trueOption, value+1, xMax, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
				acc += getAccepted(falseOption, xMin, value, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
			} else {
				acc += getAccepted(trueOption, xMin, value-1, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
				acc += getAccepted(falseOption, value, xMax, mMin, mMax, aMin, aMax, sMin, sMax, wfs)
			}
		case 'm':
			if symbol == '>' {
				acc += getAccepted(trueOption, xMin, xMax, value+1, mMax, aMin, aMax, sMin, sMax, wfs)
				acc += getAccepted(falseOption, xMin, xMax, mMin, value, aMin, aMax, sMin, sMax, wfs)
			} else {
				acc += getAccepted(trueOption, xMin, xMax, mMin, value-1, aMin, aMax, sMin, sMax, wfs)
				acc += getAccepted(falseOption, xMin, xMax, value, mMax, aMin, aMax, sMin, sMax, wfs)
			}
		case 'a':
			if symbol == '>' {
				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, value+1, aMax, sMin, sMax, wfs)
				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, value, sMin, sMax, wfs)
			} else {
				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, value-1, sMin, sMax, wfs)
				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, value, aMax, sMin, sMax, wfs)
			}
		case 's':
			if symbol == '>' {
				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, aMax, value+1, sMax, wfs)
				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value, wfs)
			} else {
				acc += getAccepted(trueOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value-1, wfs)
				acc += getAccepted(falseOption, xMin, xMax, mMin, mMax, aMin, aMax, value, sMax, wfs)
			}
		}
		return acc
	}
}

func getAllCombinations(input []string) (int, error) {
	if wfs, err := parseWorkFlows(input); err != nil {
		return 0, err
	} else {
		return getAccepted("in", 1, 4000, 1, 4000, 1, 4000, 1, 4000, wfs), nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day19/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumRatingNumbers(output))
	fmt.Println(getAllCombinations(output))
}
