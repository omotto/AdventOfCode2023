package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
	/*"math"
	"path/filepath"

	"advent2023/pkg/file"*/)

type coordinates struct {
	x, y int
}

func getMapFormInput(s []string) (m [][]uint8) {
	m = make([][]uint8, len(s))
	for y := 0; y < len(s); y++ {
		m[y] = make([]uint8, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			m[y][x] = s[y][x]
		}
	}
	return m
}

func getEmptyLines(m [][]uint8) (hLines, vLines []int) {
	var hEmpty, vEmpty bool
	for y := 0; y < len(m); y++ {
		hEmpty = true
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] != '.' {
				hEmpty = false
				break
			}
		}
		if hEmpty {
			hLines = append(hLines, y)
		}
	}
	for x := 0; x < len(m[0]); x++ {
		vEmpty = true
		for y := 0; y < len(m); y++ {
			if m[y][x] != '.' {
				vEmpty = false
				break
			}
		}
		if vEmpty {
			vLines = append(vLines, x)
		}
	}
	return hLines, vLines
}

func getCoordinates(m [][]uint8) (c []coordinates) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] != '.' {
				c = append(c, coordinates{
					x: x,
					y: y,
				})
			}
		}
	}
	return c
}

func getDistancePairs(c []coordinates, hLines, vLines []int, factor int) (p [][]int) {
	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			xDiff := c[i].x - c[j].x
			if xDiff < 0 {
				xDiff *= -1
			}
			yDiff := c[i].y - c[j].y
			if yDiff < 0 {
				yDiff *= -1
			}
			offsetY := 0
			for _, hLine := range hLines {
				if (hLine < c[i].y && hLine > c[j].y) || (hLine > c[i].y && hLine < c[j].y) {
					offsetY++
				}
			}
			offsetX := 0
			for _, vLine := range vLines {
				if (vLine < c[i].x && vLine > c[j].x) || (vLine > c[i].x && vLine < c[j].x) {
					offsetX++
				}
			}
			p = append(p, []int{i, j, xDiff + yDiff + offsetY*(factor-1) + offsetX*(factor-1)})
		}
	}
	return p
}

func getSumLengths(s []string, factor int) (sum int) {
	m := getMapFormInput(s)
	hLines, vLines := getEmptyLines(m)
	c := getCoordinates(m)                           // []coordinates{x, y}
	l := getDistancePairs(c, hLines, vLines, factor) // [][3]int{cA, cB, length}
	for _, v := range l {
		sum += v[2]
	}
	return sum
}

func main() {
	abs, _ := filepath.Abs("src/day11/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumLengths(output, 2)) // 10490062
	fmt.Println(getSumLengths(output, 1000000))
}
