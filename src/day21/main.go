package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"advent2023/pkg/file"
)

func parseInputMap(i []string) ([][]int, int, int, error) {
	sx := 0
	sy := 0
	m := make([][]int, len(i))
	for y := 0; y < len(i); y++ {
		m[y] = make([]int, len(i[y]))
		for x := 0; x < len(i[y]); x++ {
			switch i[y][x] {
			case '.':
				m[y][x] = 0
			case 'S':
				m[y][x] = 0
				sx = x
				sy = y
			case '#':
				m[y][x] = 1
			default:
				return m, sx, sy, errors.New("invalid input")
			}
		}
	}
	return m, sx, sy, nil
}

type coordinates struct {
	x, y int
}

var visited map[string]struct{}

func nextBranch(x, y, step, maxStep int, m [][]int) []coordinates {
	var acc []coordinates
	if _, ok := visited[fmt.Sprintf("%d,%d,%d", x, y, step)]; ok {
		return []coordinates{}
	} else {
		visited[fmt.Sprintf("%d,%d,%d", x, y, step)] = struct{}{}
	}
	if step == maxStep {
		return []coordinates{{x: x, y: y}}
	}
	if x > 0 && m[y][x-1] == 0 {
		acc = append(acc, nextBranch(x-1, y, step+1, maxStep, m)...)
	}
	if x < len(m[0])-1 && m[y][x+1] == 0 {
		acc = append(acc, nextBranch(x+1, y, step+1, maxStep, m)...)
	}
	if y > 0 && m[y-1][x] == 0 {
		acc = append(acc, nextBranch(x, y-1, step+1, maxStep, m)...)
	}
	if y < len(m)-1 && m[y+1][x] == 0 {
		acc = append(acc, nextBranch(x, y+1, step+1, maxStep, m)...)
	}
	return acc
}

func getGardenPlotsFilled(input []string, steps int) (int, error) {
	if m, sx, sy, err := parseInputMap(input); err != nil {
		return 0, err
	} else {
		visited = make(map[string]struct{})
		cs := nextBranch(sx, sy, 0, steps, m)
		return len(cs), nil
	}
}

func expandMap(i [][]int, factor int) [][]int {
	firstLenght := len(i)
	for y := 0; y < firstLenght; y++ {
		repeat := i[y]
		for j := 0; j < factor-1; j++ {
			i[y] = append(i[y], repeat...)
		}
	}
	for j := 0; j < factor-1; j++ {
		for y := 0; y < firstLenght; y++ {
			i = append(i, i[y])
		}
	}
	return i
}

func getGardenPlotsFilled2(input []string, steps int, factor int) (int, error) {
	if m, sx, sy, err := parseInputMap(input); err != nil {
		return 0, err
	} else {
		firstLength := len(m)
		m = expandMap(m, factor)
		sx += firstLength * (factor / 2)
		sy += firstLength * (factor / 2)
		visited = make(map[string]struct{})
		return len(nextBranch(sx, sy, 0, steps, m)), nil // 3606
	}
}

func main() {
	abs, _ := filepath.Abs("src/day21/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getGardenPlotsFilled(output, 64))
	//fmt.Println(getGardenPlotsFilled2(output, 26501365, ...))
	steps := 26501365
	p := make([]int, 3)
	size := len(output) // NxN
	half := size / 2
	p[0], _ = getGardenPlotsFilled2(output, half, 5)
	p[1], _ = getGardenPlotsFilled2(output, half+size, 5)
	p[2], _ = getGardenPlotsFilled2(output, half+2*size, 5)
	a := (p[2] + p[0] - 2*p[1]) / 2
	b := p[1] - p[0] - a
	c := p[0]
	n := steps / size
	result := a*n*n + b*n + c
	fmt.Println(result) // 584211423220706
}
