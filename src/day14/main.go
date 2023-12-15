package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"advent2023/pkg/file"
)

func parseInput(s []string) (table [][]int, err error) {
	table = make([][]int, len(s))
	for y, line := range s {
		table[y] = make([]int, len(line))
		for x := 0; x < len(line); x++ {
			switch line[x] {
			case 'O':
				table[y][x] = 2
			case '#':
				table[y][x] = 1
			case '.':
				table[y][x] = 0
			default:
				return table, errors.New("invalid input")
			}
		}
	}
	return table, err
}

func getTotalLoad(s []string) (int, error) {
	var result int
	if table, err := parseInput(s); err != nil {
		return 0, err
	} else {
		table = tiltNorth(table)
		var numRocks int
		for y := 0; y < len(table); y++ {
			numRocks = 0
			for x := 0; x < len(table[y]); x++ {
				if table[y][x] == 2 {
					numRocks++
				}
			}
			result += numRocks * (len(table) - y)
		}
	}
	return result, nil
}

func tiltNorth(table [][]int) [][]int {
	var stopPoint int
	for x := 0; x < len(table[0]); x++ {
		stopPoint = 0
		for y := 0; y < len(table); y++ {
			if table[y][x] == 1 {
				stopPoint = y + 1
			}
			if table[y][x] == 2 {
				if stopPoint == y {
					stopPoint = y + 1
				} else {
					table[stopPoint][x] = 2
					table[y][x] = 0
					stopPoint = stopPoint + 1
				}
			}
		}
	}
	return table
}

func tiltSouth(table [][]int) [][]int {
	var stopPoint int
	for x := 0; x < len(table[0]); x++ {
		stopPoint = len(table) - 1
		for y := len(table) - 1; y >= 0; y-- {
			if table[y][x] == 1 {
				stopPoint = y - 1
			}
			if table[y][x] == 2 {
				if stopPoint == y {
					stopPoint = y - 1
				} else {
					table[stopPoint][x] = 2
					table[y][x] = 0
					stopPoint = stopPoint - 1
				}
			}
		}
	}
	return table
}

func tiltWest(table [][]int) [][]int {
	var stopPoint int
	for y := 0; y < len(table); y++ {
		stopPoint = 0
		for x := 0; x < len(table[0]); x++ {
			if table[y][x] == 1 {
				stopPoint = x + 1
			}
			if table[y][x] == 2 {
				if stopPoint == x {
					stopPoint = x + 1
				} else {
					table[y][stopPoint] = 2
					table[y][x] = 0
					stopPoint = stopPoint + 1
				}
			}
		}
	}
	return table
}

func tiltEast(table [][]int) [][]int {
	var stopPoint int
	for y := 0; y < len(table); y++ {
		stopPoint = len(table[0]) - 1
		for x := len(table[0]) - 1; x >= 0; x-- {
			if table[y][x] == 1 {
				stopPoint = x - 1
			}
			if table[y][x] == 2 {
				if stopPoint == x {
					stopPoint = x - 1
				} else {
					table[y][stopPoint] = 2
					table[y][x] = 0
					stopPoint = stopPoint - 1
				}
			}
		}
	}
	return table
}

func getTotalLoad2(s []string, numCycles int) (int, error) {
	var result int
	if table, err := parseInput(s); err != nil {
		return 0, err
	} else {
		for i := 0; i < numCycles; i++ {
			table = tiltNorth(table)
			table = tiltWest(table)
			table = tiltSouth(table)
			table = tiltEast(table)
		}
		var numRocks int
		for y := 0; y < len(table); y++ {
			numRocks = 0
			for x := 0; x < len(table[y]); x++ {
				if table[y][x] == 2 {
					numRocks++
				}
			}
			result += numRocks * (len(table) - y)
		}
	}
	return result, nil
}

func getSequence(values []int) (frequency, offset int, ok bool) {
	if len(values) > 10 {
		var reps []int
		for i := 0; i < len(values)-1; i++ {
			if values[i] == values[len(values)-1] {
				reps = append(reps, i)
				if len(reps) == 3 {
					if reps[2]-reps[1] == reps[1]-reps[0] && reps[2]-reps[1] > 1 {
						ok = true
						for i := 0; i < reps[1]-reps[0]; i++ {
							if values[reps[0]+i] != values[reps[1]+i] {
								ok = false
							}
						}
						if ok {
							frequency = reps[2] - reps[1]
							offset = reps[0]
						}
					}
				}
			}
		}
	}
	return frequency, offset, ok
}

func getTotalLoad3(s []string, numCycles int) (int, error) {
	var result int
	if table, err := parseInput(s); err != nil {
		return 0, err
	} else {
		var results []int
		for {
			table = tiltNorth(table)
			table = tiltWest(table)
			table = tiltSouth(table)
			table = tiltEast(table)
			numRocks := 0
			result = 0
			for y := 0; y < len(table); y++ {
				numRocks = 0
				for x := 0; x < len(table[y]); x++ {
					if table[y][x] > 1 {
						numRocks++
					}
				}
				result += numRocks * (len(table) - y)
			}
			results = append(results, result)
			if frequency, offset, ok := getSequence(results); ok {
				mod := (numCycles - offset) % frequency
				result = results[offset+mod-1]
				break
			}
		}
	}
	return result, nil
}

func main() {
	abs, _ := filepath.Abs("src/day14/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getTotalLoad(output))
	fmt.Println(getTotalLoad2(output, 10000))      // 79723
	fmt.Println(getTotalLoad3(output, 1000000000)) // 79723
}
