package main

import (
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

func parseLine(s string) (direction uint8, numTiles int, err error) {
	r := strings.Split(s, " ")
	if len(r) != 3 {
		return direction, numTiles, fmt.Errorf("invalid input")
	}
	if len(r[0]) != 1 {
		return direction, numTiles, fmt.Errorf("invalid input")
	}
	direction = r[0][0]
	if direction != 'U' && direction != 'D' && direction != 'L' && direction != 'R' {
		return direction, numTiles, fmt.Errorf("invalid input")
	}
	numTiles, err = strconv.Atoi(r[1])
	return direction, numTiles, err
}

func createEmptyMap(instructions []string) ([][]int, int, int, error) {
	var (
		up, down, left, right int
		m                     [][]int
	)
	for _, instruction := range instructions {
		if direction, numTiles, err := parseLine(instruction); err != nil {
			return m, 0, 0, err
		} else {
			switch direction {
			case 'U':
				up += numTiles
			case 'D':
				down += numTiles
			case 'L':
				left += numTiles
			case 'R':
				right += numTiles
			}
		}
	}
	m = make([][]int, up+down)
	for y := 0; y < len(m); y++ {
		m[y] = make([]int, right+left)
	}
	return m, up - 1, right - 1, nil
}

func createOutline(instructions []string) (m [][]int, err error) {
	var sx, sy int
	m, sy, sx, err = createEmptyMap(instructions)
	if err != nil {
		return nil, err
	}
	for _, instruction := range instructions {
		direction, numTiles, _ := parseLine(instruction)
		switch direction {
		case 'U':
			for y := 0; y < numTiles; y++ {
				m[sy-y][sx] = 1
			}
			sy = sy - numTiles
		case 'D':
			for y := 0; y < numTiles; y++ {
				m[sy+y][sx] = 1
			}
			sy = sy + numTiles
		case 'L':
			for x := 0; x < numTiles; x++ {
				m[sy][sx-x] = 1
			}
			sx = sx - numTiles
		case 'R':
			for x := 0; x < numTiles; x++ {
				m[sy][sx+x] = 1
			}
			sx = sx + numTiles
		}
	}
	return m, err
}

func fillMap(inputMap [][]int) ([][]int, int) {
	var (
		hMap    [][]int
		total   int
		paint   bool
		started uint8
	)
	hMap = make([][]int, len(inputMap))
	for y := 0; y < len(inputMap); y++ {
		hMap[y] = make([]int, len(inputMap[y]))
		paint = false
		for x := 0; x < len(inputMap[y]); x++ {
			if x > 0 && x < len(inputMap[y])-1 && y > 0 && y < len(inputMap)-1 {
				if inputMap[y][x] == 1 {
					if inputMap[y][x-1] == 0 && inputMap[y][x+1] == 0 {
						paint = !paint
					} else if inputMap[y][x+1] == 1 && inputMap[y][x-1] == 0 {
						if inputMap[y+1][x] == 1 {
							started = 'F'
						}
						if inputMap[y-1][x] == 1 {
							started = 'L'
						}
					} else if inputMap[y][x+1] == 0 && inputMap[y][x-1] == 1 {
						if inputMap[y+1][x] == 1 && started == 'L' {
							paint = !paint
						}
						if inputMap[y-1][x] == 1 && started == 'F' {
							paint = !paint
						}
					}
				}
			}
			if paint {
				hMap[y][x] = 1
			}
		}
	}
	total = 0
	for y := 0; y < len(inputMap); y++ {
		for x := 0; x < len(inputMap[y]); x++ {
			if hMap[y][x] == 1 {
				inputMap[y][x] = 1
			}
			total += inputMap[y][x]
		}
	}
	return inputMap, total
}

func getNumFilledTiles(s []string) (int, error) {
	var sum int = 0
	if m, err := createOutline(s); err != nil {
		return sum, err
	} else {
		m, sum = fillMap(m)
	}
	return sum, nil
}

func parseLine2(s string) (direction uint8, numTiles int, err error) {
	r := strings.Split(s, " ")
	if len(r) != 3 {
		return direction, numTiles, fmt.Errorf("invalid input")
	}
	numberStr := r[2][2 : len(r[2])-2]
	n := new(big.Int)
	n.SetString(numberStr, 16)
	numTiles = int(n.Int64())
	switch r[2][len(r[2])-2] {
	case '0':
		direction = 'R'
	case '1':
		direction = 'D'
	case '2':
		direction = 'L'
	case '3':
		direction = 'U'
	default:
		return direction, numTiles, fmt.Errorf("invalid input")
	}
	return direction, numTiles, err
}

type corner struct {
	x, y int
}

func createMap(i []string) ([]corner, int, error) {
	m := make([]corner, len(i))
	sx := 0
	sy := 0
	outline := 0
	for idx, instruction := range i {
		if direction, numTiles, err := parseLine2(instruction); err != nil {
			return m, 0, err
		} else {
			m[idx] = corner{
				x: sx,
				y: sy,
			}
			switch direction {
			case 'U':
				sy = sy - numTiles
			case 'D':
				sy = sy + numTiles
			case 'L':
				sx = sx - numTiles
			case 'R':
				sx = sx + numTiles
			}
			outline += numTiles
		}
	}
	return m, outline, nil
}

func shoelace(m []corner) int {
	var a, b int
	for i := 0; i < len(m)-1; i++ {
		a += m[i].x * m[i+1].y
		b += m[i].y * m[i+1].x
	}
	c := a - b
	if c < 0 {
		c = -c
	}
	return c / 2
}

func getNumFilledTiles2(s []string) (int, error) {
	if m, outline, err := createMap(s); err != nil {
		return 0, err
	} else {
		return shoelace(m) + outline/2 + 1, nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day18/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getNumFilledTiles(output))
	fmt.Println(getNumFilledTiles2(output))
}
