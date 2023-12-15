package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type cube struct {
	blue  int
	red   int
	green int
}

const (
	game  = "game"
	green = "green"
	red   = "red"
	blue  = "blue"
)

func parseInput(s string) (id int, values []cube, err error) {
	r := strings.Split(s, ": ")
	if len(r) != 2 {
		return 0, nil, errors.New("invalid input string")
	}
	r[0] = strings.ToLower(r[0])
	r0 := strings.Split(r[0], " ")
	if len(r0) != 2 {
		return 0, nil, errors.New("invalid input string")
	}
	if r0[0] != game {
		return 0, nil, errors.New("invalid input string")
	}
	id, err = strconv.Atoi(r0[1])
	if err != nil {
		return 0, nil, err
	}
	r[1] = strings.ToLower(r[1])
	r1 := strings.Split(r[1], "; ")
	for _, g := range r1 {
		v := strings.Split(g, ", ")
		var cubes cube
		for _, vv := range v {
			f := strings.Split(vv, " ")
			if len(f) != 2 {
				return 0, nil, errors.New("invalid input string")
			}
			switch f[1] {
			case blue:
				cubes.blue, err = strconv.Atoi(f[0])
			case green:
				cubes.green, err = strconv.Atoi(f[0])
			case red:
				cubes.red, err = strconv.Atoi(f[0])
			default:
				return 0, nil, errors.New("invalid input string")
			}
			if err != nil {
				return 0, nil, err
			}
		}
		values = append(values, cubes)
	}
	return id, values, err
}

func getMinimumSetOfCubes(games []cube) cube {
	var (
		maxR = 0
		maxG = 0
		maxB = 0
	)
	for _, g := range games {
		if g.red > maxR {
			maxR = g.red
		}
		if g.blue > maxB {
			maxB = g.blue
		}
		if g.green > maxG {
			maxG = g.green
		}
	}
	return cube{
		blue:  maxB,
		red:   maxR,
		green: maxG,
	}
}

func checkValidGame(games []cube, maxG, maxR, maxB int) (ok bool) {
	for _, g := range games {
		if g.red > maxR {
			return false
		}
		if g.blue > maxB {
			return false
		}
		if g.green > maxG {
			return false
		}
	}
	return true
}

func getSumValidIDGames(s []string) (sum int, err error) {
	var (
		id     int
		values []cube
	)
	for _, t := range s {
		id, values, err = parseInput(t)
		if err != nil {
			return 0, err
		}
		if valid := checkValidGame(values, 13, 12, 14); valid {
			sum += id
		}
	}
	return sum, err
}

func getSumPoweredCubes(s []string) (sum int, err error) {
	var values []cube
	for _, t := range s {
		_, values, err = parseInput(t)
		if err != nil {
			return 0, err
		}
		minimum := getMinimumSetOfCubes(values)
		sum += minimum.green * minimum.red * minimum.blue
	}
	return sum, err
}

func main() {
	abs, _ := filepath.Abs("src/day02/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumValidIDGames(output))
	fmt.Println(getSumPoweredCubes(output))
}
