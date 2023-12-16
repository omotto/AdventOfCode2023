package main

import (
	"fmt"
	"path/filepath"

	"advent2023/pkg/file"
)

const (
	Right int = iota
	Left
	Up
	Down
)

type position struct {
	x, y, direction int
}

func calculateTableRecursive(x, y int, m []string, table [][]int, direction int) {
	var mask int
	xx := x
	yy := y
	for {
		if yy < 0 || yy > len(m)-1 || xx > len(m[yy])-1 || xx < 0 {
			break
		}
		switch direction {
		case Right:
			mask = 1
		case Left:
			mask = 2
		case Up:
			mask = 4
		case Down:
			mask = 8
		}
		if table[yy][xx]&mask > 0 { // visited tile
			break
		}
		table[yy][xx] |= mask
		switch m[yy][xx] {
		case '.':
			switch direction {
			case Right:
				xx++
			case Left:
				xx--
			case Up:
				yy--
			case Down:
				yy++
			}
		case '\\':
			switch direction {
			case Right:
				yy++
				direction = Down
			case Left:
				yy--
				direction = Up
			case Up:
				xx--
				direction = Left
			case Down:
				xx++
				direction = Right
			}
		case '/':
			switch direction {
			case Right:
				yy--
				direction = Up
			case Left:
				yy++
				direction = Down
			case Up:
				xx++
				direction = Right
			case Down:
				xx--
				direction = Left
			}
		case '-':
			switch direction {
			case Right:
				xx++
			case Left:
				xx--
			case Up, Down:
				calculateTableRecursive(xx-1, yy, m, table, Left)
				calculateTableRecursive(xx+1, yy, m, table, Right)
				yy = -1 // Force exit
				break
			}
		case '|':
			switch direction {
			case Right, Left:
				calculateTableRecursive(xx, yy-1, m, table, Up)
				calculateTableRecursive(xx, yy+1, m, table, Down)
				xx = -1 // Force exit
			case Up:
				yy--
			case Down:
				yy++
			}
		}
	}
}

func calculateTable(x, y int, m []string, table [][]int, firstDirection int) {
	var mask int
	positions := []position{
		{
			x:         x,
			y:         y,
			direction: firstDirection,
		},
	}
	for {
		if len(positions) == 0 {
			break // if stack is empty goes out
		}
		p := positions[len(positions)-1]         // get last
		positions = positions[:len(positions)-1] // remove the last one
		xx := p.x
		yy := p.y
		direction := p.direction
		for {
			if yy < 0 || yy > len(m)-1 || xx > len(m[yy])-1 || xx < 0 { // out of map
				break
			}
			switch direction {
			case Right:
				mask = 1
			case Left:
				mask = 2
			case Up:
				mask = 4
			case Down:
				mask = 8
			}
			if table[yy][xx]&mask > 0 { // visited tile
				break
			}
			table[yy][xx] |= mask
			switch m[yy][xx] {
			case '.':
				switch direction {
				case Right:
					xx++
				case Left:
					xx--
				case Up:
					yy--
				case Down:
					yy++
				}
			case '\\':
				switch direction {
				case Right:
					yy++
					direction = Down
				case Left:
					yy--
					direction = Up
				case Up:
					xx--
					direction = Left
				case Down:
					xx++
					direction = Right
				}
			case '/':
				switch direction {
				case Right:
					yy--
					direction = Up
				case Left:
					yy++
					direction = Down
				case Up:
					xx++
					direction = Right
				case Down:
					xx--
					direction = Left
				}
			case '-':
				switch direction {
				case Right:
					xx++
				case Left:
					xx--
				case Up, Down:
					positions = append(positions, position{
						x:         xx - 1,
						y:         yy,
						direction: Left,
					})
					positions = append(positions, position{
						x:         xx + 1,
						y:         yy,
						direction: Right,
					})
					yy = -1 // Force exit
				}
			case '|':
				switch direction {
				case Right, Left:
					positions = append(positions, position{
						x:         xx,
						y:         yy - 1,
						direction: Up,
					})
					positions = append(positions, position{
						x:         xx,
						y:         yy + 1,
						direction: Down,
					})
					xx = -1 // Force exit
				case Up:
					yy--
				case Down:
					yy++
				}
			}
		}
	}
}

func getSumTilesEnergized(s []string, xx, yy, direction int) int {
	sum := 0
	table := make([][]int, len(s))
	for i, _ := range table {
		table[i] = make([]int, len(s[i]))
	}
	calculateTable(xx, yy, s, table, direction)
	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] != 0 {
				sum++
			}
		}
	}
	return sum
}

func getSumTilesMaxEnergized(s []string) int {
	max := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			sum := 0
			if y == 0 {
				sum += getSumTilesEnergized(s, x, y, Down)
			}
			if x == 0 {
				sum += getSumTilesEnergized(s, x, y, Right)
			}
			if y == len(s)-1 {
				sum += getSumTilesEnergized(s, x, y, Up)
			}
			if x == len(s[y])-1 {
				sum += getSumTilesEnergized(s, x, y, Left)
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	abs, _ := filepath.Abs("src/day16/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumTilesEnergized(output, 0, 0, Right))
	fmt.Println(getSumTilesMaxEnergized(output))
}
