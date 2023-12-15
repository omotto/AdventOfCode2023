package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func getMapDistance(s []string, m [][]int, level int) [][]int {
	nextLevel := level + 1
	found := false
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if m[y][x] == level {
				found = true
				switch s[y][x] {
				case 'S':
					if x > 0 && m[y][x-1] == -1 && (s[y][x-1] == '-' || s[y][x-1] == 'L' || s[y][x-1] == 'F') {
						m[y][x-1] = nextLevel
					}
					if x < len(m[y])-1 && m[y][x+1] == -1 && (s[y][x+1] == '-' || s[y][x+1] == '7' || s[y][x+1] == 'J') {
						m[y][x+1] = nextLevel
					}
					if y > 0 && m[y-1][x] == -1 && (s[y-1][x] == '|' || s[y-1][x] == '7' || s[y-1][x] == 'F') {
						m[y-1][x] = nextLevel
					}
					if y < len(m)-1 && m[y+1][x] == -1 && (s[y+1][x] == '|' || s[y+1][x] == 'J' || s[y+1][x] == 'L') {
						m[y+1][x] = nextLevel
					}
				case '-':
					if x > 0 && m[y][x-1] == -1 && (s[y][x-1] == '-' || s[y][x-1] == 'L' || s[y][x-1] == 'F') {
						m[y][x-1] = nextLevel
					}
					if x < len(m[y])-1 && m[y][x+1] == -1 && (s[y][x+1] == '-' || s[y][x+1] == '7' || s[y][x+1] == 'J') {
						m[y][x+1] = nextLevel
					}
				case '|':
					if y > 0 && m[y-1][x] == -1 && (s[y-1][x] == '|' || s[y-1][x] == '7' || s[y-1][x] == 'F') {
						m[y-1][x] = nextLevel
					}
					if y < len(m)-1 && m[y+1][x] == -1 && (s[y+1][x] == '|' || s[y+1][x] == 'J' || s[y+1][x] == 'L') {
						m[y+1][x] = nextLevel
					}
				case '7':
					if x > 0 && m[y][x-1] == -1 && (s[y][x-1] == '-' || s[y][x-1] == 'L' || s[y][x-1] == 'F') {
						m[y][x-1] = nextLevel
					}
					if y < len(m)-1 && m[y+1][x] == -1 && (s[y+1][x] == '|' || s[y+1][x] == 'J' || s[y+1][x] == 'L') {
						m[y+1][x] = nextLevel
					}
				case 'F':
					if x < len(m[y])-1 && m[y][x+1] == -1 && (s[y][x+1] == '-' || s[y][x+1] == '7' || s[y][x+1] == 'J') {
						m[y][x+1] = nextLevel
					}
					if y < len(m)-1 && m[y+1][x] == -1 && (s[y+1][x] == '|' || s[y+1][x] == 'J' || s[y+1][x] == 'L') {
						m[y+1][x] = nextLevel
					}
				case 'J':
					if x > 0 && m[y][x-1] == -1 && (s[y][x-1] == '-' || s[y][x-1] == 'L' || s[y][x-1] == 'F') {
						m[y][x-1] = nextLevel
					}
					if y > 0 && m[y-1][x] == -1 && (s[y-1][x] == '|' || s[y-1][x] == '7' || s[y-1][x] == 'F') {
						m[y-1][x] = nextLevel
					}
				case 'L':
					if x < len(m[y])-1 && m[y][x+1] == -1 && (s[y][x+1] == '-' || s[y][x+1] == '7' || s[y][x+1] == 'J') {
						m[y][x+1] = nextLevel
					}
					if y > 0 && m[y-1][x] == -1 && (s[y-1][x] == '|' || s[y-1][x] == '7' || s[y-1][x] == 'F') {
						m[y-1][x] = nextLevel
					}
				}
			}
		}
	}
	if found {
		return getMapDistance(s, m, nextLevel)
	}
	return m
}

func getFarthestStep(s []string) int {
	m := make([][]int, len(s))
	for y := 0; y < len(s); y++ {
		m[y] = make([]int, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == 'S' {
				m[y][x] = 0
			} else if s[y][x] == '.' {
				m[y][x] = -2
			} else {
				m[y][x] = -1
			}
		}
	}
	m = getMapDistance(s, m, 0)
	var max int = 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if m[y][x] > max {
				max = m[y][x]
			}
		}
	}
	return max
}

func getStartSymbol(s []string, sy, sx int) uint8 {
	var (
		up    uint8 = '.'
		down  uint8 = '.'
		left  uint8 = '.'
		right uint8 = '.'
	)
	if sy > 0 {
		up = s[sy-1][sx]
	}
	if sy < len(s)-1 {
		down = s[sy+1][sx]
	}
	if sx > 0 {
		left = s[sy][sx-1]
	}
	if sx < len(s[0])-1 {
		right = s[sy][sx+1]
	}
	upCond := up == '7' || up == 'F' || up == '|'
	downCond := down == 'J' || down == 'L' || down == '|'
	leftCond := left == 'F' || left == 'L' || left == '-'
	rightCond := right == 'J' || right == '7' || right == '-'

	if upCond && downCond {
		return '|'
	}
	if leftCond && rightCond {
		return '-'
	}
	if leftCond && downCond {
		return '7'
	}
	if rightCond && downCond {
		return 'F'
	}
	if rightCond && upCond {
		return 'L'
	}
	if leftCond && upCond {
		return 'J'
	}
	return 'S'
}

func getTilesEnclosed(s []string) int {
	var sy, sx int
	m := make([][]int, len(s))
	for y := 0; y < len(s); y++ {
		m[y] = make([]int, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == 'S' {
				m[y][x] = 0
				sy = y
				sx = x
			} else if s[y][x] == '.' {
				m[y][x] = -2
			} else {
				m[y][x] = -1
			}
		}
	}
	m = getMapDistance(s, m, 0)
	s[sy] = strings.Replace(s[sy], "S", string(getStartSymbol(s, sy, sx)), -1)
	for y := 0; y < len(s); y++ {
		newS := make([]byte, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if m[y][x] < 0 {
				newS[x] = '.'
			} else {
				newS[x] = s[y][x]
			}
		}
		s[y] = string(newS)
	}
	noWall := regexp.MustCompile(`F-*7|L-*J`)
	wall := regexp.MustCompile(`F-*J|L-*7`)
	for i, ss := range s {
		s1 := noWall.ReplaceAllString(ss, " ")
		s2 := wall.ReplaceAllString(s1, "|")
		s[i] = s2
	}
	var (
		parity int
		count  int = 0
	)
	for y := 0; y < len(s); y++ {
		parity = 0
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '|' {
				parity++
			}
			if s[y][x] == '.' && parity%2 == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	abs, _ := filepath.Abs("src/day10/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getFarthestStep(output))
	fmt.Println(getTilesEnclosed(output))
}
