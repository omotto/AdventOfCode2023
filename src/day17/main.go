package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
)

func distance(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	dy := y1 - y2
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func findShortestPathRecursive(s []string, sx, sy, dx, dy, up, down, left, right, value int, visited map[string]int, final []int) []int {
	var ret []int
	iVal, _ := strconv.Atoi(string(s[sy][sx]))
	weight := value + iVal
	if sx == dx && sy == dy {
		return []int{weight}
	}
	key := fmt.Sprintf("%d,%d,%d,%d,%d,%d", sx, sy, up, down, left, right)
	if v, ok := visited[key]; !ok || ok && v > weight {
		visited[key] = weight
		for _, w := range final {
			if weight > w-distance(sx, sy, dx, dy) {
				return ret
			}
		}
		if sx > 0 && left < 3 && right == 0 {
			ret = append(ret, findShortestPathRecursive(s, sx-1, sy, dx, dy, 0, 0, left+1, 0, weight, visited, ret)...)
		}
		if sx < len(s[0])-1 && right < 3 && left == 0 {
			ret = append(ret, findShortestPathRecursive(s, sx+1, sy, dx, dy, 0, 0, 0, right+1, weight, visited, ret)...)
		}
		if sy > 0 && up < 3 && down == 0 {
			ret = append(ret, findShortestPathRecursive(s, sx, sy-1, dx, dy, up+1, 0, 0, 0, weight, visited, ret)...)
		}
		if sy < len(s)-1 && down < 3 && up == 0 {
			ret = append(ret, findShortestPathRecursive(s, sx, sy+1, dx, dy, 0, down+1, 0, 0, weight, visited, ret)...)
		}
	}
	return ret
}

func findShortestPathTemperature(s []string) int {
	visited := make(map[string]int)
	paths := findShortestPathRecursive(s, 0, 0, len(s[0])-1, len(s)-1, 0, 0, 0, 0, 0, visited, []int{})
	min := 9999999999
	for _, path := range paths {
		if path < min {
			min = path
		}
	}
	sVal, _ := strconv.Atoi(string(s[0][0]))
	return min - sVal
}

type position struct {
	x         int
	y         int
	accWeight int
	up        int
	down      int
	left      int
	right     int
}

func findShortestPath(s []string) int {
	var (
		priorityQueue []position
		sx            = 0
		sy            = 0
		dx            = len(s[0]) - 1
		dy            = len(s) - 1
	)
	priorityQueue = append(priorityQueue, position{
		x:     sx + 1,
		y:     0,
		up:    0,
		down:  0,
		left:  0,
		right: 1,
	})
	priorityQueue = append(priorityQueue, position{
		x:     0,
		y:     sy + 1,
		up:    0,
		down:  1,
		left:  0,
		right: 0,
	})
	visited := make(map[string]int)
	for len(priorityQueue) > 0 {
		sort.SliceStable(priorityQueue, func(i, j int) bool {
			return priorityQueue[i].accWeight < priorityQueue[j].accWeight
		})
		//
		entry := priorityQueue[0]
		if len(priorityQueue) > 1 {
			priorityQueue = priorityQueue[1:]
		} else {
			priorityQueue = []position{}
		}
		//
		iVal, _ := strconv.Atoi(string(s[entry.y][entry.x]))
		weight := entry.accWeight + iVal
		if entry.x == dx && entry.y == dy {
			return weight
		}
		//
		key := fmt.Sprintf("%d,%d,%d,%d,%d,%d", entry.x, entry.y, entry.up, entry.down, entry.left, entry.right)
		if v, ok := visited[key]; !ok || ok && v > weight {
			visited[key] = weight
			if entry.x > 0 && entry.left < 3 && entry.right == 0 {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x - 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      entry.left + 1,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.x < len(s[0])-1 && entry.right < 3 && entry.left == 0 {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x + 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      0,
					right:     entry.right + 1,
					accWeight: weight,
				})
			}
			if entry.y > 0 && entry.up < 3 && entry.down == 0 {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y - 1,
					up:        entry.up + 1,
					down:      0,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.y < len(s)-1 && entry.down < 3 && entry.up == 0 {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y + 1,
					up:        0,
					down:      entry.down + 1,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
		}
	}
	return 0
}

func findShortestPath2(s []string, minSteps, maxSteps int) int {
	var (
		priorityQueue []position
		sx            = 0
		sy            = 0
		dx            = len(s[0]) - 1
		dy            = len(s) - 1
	)
	priorityQueue = append(priorityQueue, position{
		x:     sx + 1,
		y:     0,
		up:    0,
		down:  0,
		left:  0,
		right: 1,
	})
	priorityQueue = append(priorityQueue, position{
		x:     0,
		y:     sy + 1,
		up:    0,
		down:  1,
		left:  0,
		right: 0,
	})
	visited := make(map[string]int)
	for len(priorityQueue) > 0 {
		sort.SliceStable(priorityQueue, func(i, j int) bool {
			return priorityQueue[i].accWeight < priorityQueue[j].accWeight
		})
		//
		entry := priorityQueue[0]
		if len(priorityQueue) > 1 {
			priorityQueue = priorityQueue[1:]
		} else {
			priorityQueue = []position{}
		}
		//
		iVal, _ := strconv.Atoi(string(s[entry.y][entry.x]))
		weight := entry.accWeight + iVal
		if entry.x == dx && entry.y == dy && (entry.left >= minSteps || entry.right >= minSteps || entry.up >= minSteps || entry.down >= minSteps) {
			return weight
		}
		//
		key := fmt.Sprintf("%d,%d,%d,%d,%d,%d", entry.x, entry.y, entry.up, entry.down, entry.left, entry.right)
		if v, ok := visited[key]; !ok || ok && v > weight {
			visited[key] = weight
			if entry.x > 0 && entry.right == 0 && (((entry.up >= minSteps || entry.down >= minSteps) && entry.left == 0) || entry.left > 0 && entry.left < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x - 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      entry.left + 1,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.x < len(s[0])-1 && entry.left == 0 && (((entry.up >= minSteps || entry.down >= minSteps) && entry.right == 0) || entry.right > 0 && entry.right < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x + 1,
					y:         entry.y,
					up:        0,
					down:      0,
					left:      0,
					right:     entry.right + 1,
					accWeight: weight,
				})
			}
			if entry.y > 0 && entry.down == 0 && (((entry.left >= minSteps || entry.right >= minSteps) && entry.up == 0) || entry.up > 0 && entry.up < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y - 1,
					up:        entry.up + 1,
					down:      0,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
			if entry.y < len(s)-1 && entry.up == 0 && (((entry.left >= minSteps || entry.right >= minSteps) && entry.down == 0) || entry.down > 0 && entry.down < maxSteps) {
				priorityQueue = append(priorityQueue, position{
					x:         entry.x,
					y:         entry.y + 1,
					up:        0,
					down:      entry.down + 1,
					left:      0,
					right:     0,
					accWeight: weight,
				})
			}
		}
	}
	return 0
}

func main() {
	abs, _ := filepath.Abs("src/day17/input.txt")
	output, _ := file.ReadInput(abs)
	//fmt.Println(findShortestPathTemperature(output))
	fmt.Println(findShortestPath(output))
	fmt.Println(findShortestPath2(output, 4, 10))
}
