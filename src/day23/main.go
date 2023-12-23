package main

import (
	"advent2023/pkg/file"
	"errors"
	"fmt"
	"path/filepath"
)

type node struct {
	x, y    int
	step    int
	visited map[string]struct{}
}

func getInitialPoint(m []string) (int, int, error) {
	if len(m) < 1 {
		return -1, -1, errors.New("invalid input")
	}
	y := 0
	for x := 0; x < len(m); x++ {
		if m[y][x] == '.' {
			return x, y, nil
		}
	}
	return -1, -1, errors.New("invalid input")
}

func getFinalPoint(m []string) (int, int, error) {
	if len(m) < 1 {
		return -1, -1, errors.New("invalid input")
	}
	y := len(m) - 1
	for x := 0; x < len(m); x++ {
		if m[y][x] == '.' {
			return x, y, nil
		}
	}
	return -1, -1, errors.New("invalid input")
}

func getLongestPathTiles(m []string, slopes bool) (int, error) {
	var steps []int
	if sx, sy, e1 := getInitialPoint(m); e1 != nil {
		return -1, e1
	} else if dx, dy, e2 := getFinalPoint(m); e2 != nil {
		return -1, e2
	} else {
		queue := []node{
			{
				x:       sx,
				y:       sy,
				step:    0,
				visited: map[string]struct{}{fmt.Sprintf("%d,%d", sx, sy): {}},
			},
		}
		for {
			// exit condition
			if len(queue) == 0 {
				break
			}
			// pop
			n := queue[0]
			queue = queue[1:]
			// check final node
			if n.x == dx && n.y == dy {
				steps = append(steps, n.step)
			} else {
				// check UP
				if _, isVisited := n.visited[fmt.Sprintf("%d,%d", n.x, n.y-1)]; !isVisited {
					if (n.y > 0 && m[n.y-1][n.x] != '#' && m[n.y-1][n.x] != 'v' && (m[n.y][n.x] == '.' || m[n.y][n.x] == '^') && slopes) ||
						(n.y > 0 && m[n.y-1][n.x] != '#' && !slopes) {
						visited := make(map[string]struct{})
						for key, value := range n.visited {
							visited[key] = value
						}
						visited[fmt.Sprintf("%d,%d", n.x, n.y-1)] = struct{}{}
						queue = append(queue, node{
							x:       n.x,
							y:       n.y - 1,
							step:    n.step + 1,
							visited: visited,
						})
					}
				}
				// check DOWN
				if _, isVisited := n.visited[fmt.Sprintf("%d,%d", n.x, n.y+1)]; !isVisited {
					if (n.y < len(m)-1 && m[n.y+1][n.x] != '#' && m[n.y+1][n.x] != '^' && (m[n.y][n.x] == '.' || m[n.y][n.x] == 'v') && slopes) ||
						(n.y < len(m)-1 && m[n.y+1][n.x] != '#' && !slopes) {
						visited := make(map[string]struct{})
						for key, value := range n.visited {
							visited[key] = value
						}
						visited[fmt.Sprintf("%d,%d", n.x, n.y+1)] = struct{}{}
						queue = append(queue, node{
							x:       n.x,
							y:       n.y + 1,
							step:    n.step + 1,
							visited: visited,
						})
					}
				}
				// check LEFT
				if _, isVisited := n.visited[fmt.Sprintf("%d,%d", n.x-1, n.y)]; !isVisited {
					if (n.x > 0 && m[n.y][n.x-1] != '#' && m[n.y][n.x-1] != '>' && (m[n.y][n.x] == '.' || m[n.y][n.x] == '<') && slopes) ||
						(n.x > 0 && m[n.y][n.x-1] != '#' && !slopes) {
						visited := make(map[string]struct{})
						for key, value := range n.visited {
							visited[key] = value
						}
						visited[fmt.Sprintf("%d,%d", n.x-1, n.y)] = struct{}{}
						queue = append(queue, node{
							x:       n.x - 1,
							y:       n.y,
							step:    n.step + 1,
							visited: visited,
						})
					}
				}
				// check RIGHT
				if _, isVisited := n.visited[fmt.Sprintf("%d,%d", n.x+1, n.y)]; !isVisited {
					if (n.x < len(m[0])-1 && m[n.y][n.x+1] != '#' && m[n.y][n.x+1] != '<' && (m[n.y][n.x] == '.' || m[n.y][n.x] == '>') && slopes) ||
						(n.x < len(m[0])-1 && m[n.y][n.x+1] != '#' && !slopes) {
						visited := make(map[string]struct{})
						for key, value := range n.visited {
							visited[key] = value
						}
						visited[fmt.Sprintf("%d,%d", n.x+1, n.y)] = struct{}{}
						queue = append(queue, node{
							x:       n.x + 1,
							y:       n.y,
							step:    n.step + 1,
							visited: visited,
						})
					}
				}
			}
		}
	}
	maxSteps := 0
	for _, step := range steps {
		if step > maxSteps {
			maxSteps = step
		}
	}
	return maxSteps, nil
}

type node2 struct {
	x, y      int
	step      int
	direction int
	visited   map[string]struct{}
}

func getLongestPathTiles2(m []string) (int, error) {
	maxSteps := 0
	if sx, sy, e1 := getInitialPoint(m); e1 != nil {
		return -1, e1
	} else if dx, dy, e2 := getFinalPoint(m); e2 != nil {
		return -1, e2
	} else {
		deadQueue := make(map[string]struct{})
		queue := []node2{
			{
				x:         sx,
				y:         sy,
				step:      0,
				direction: 1,
				visited:   map[string]struct{}{fmt.Sprintf("%d,%d", sx, sy): {}},
			},
		}
		for {
			// exit condition
			if len(queue) == 0 {
				break
			}
			// pop
			n := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			// search next node
			x := n.x
			y := n.y
			step := n.step
			nextDirection := n.direction
			for {
				switch nextDirection {
				case 0: //  UP
					y--
				case 1: //  DOWN
					y++
				case 2: //  LEFT
					x--
				case 3: //  RIGHT
					x++
				}
				step++
				// check final node
				if x == dx && y == dy {
					if step > maxSteps {
						maxSteps = step
					}
					break
				}
				directions := make([]int, 4)
				// check UP
				if y > 0 && m[y-1][x] != '#' && nextDirection != 1 {
					directions[0] = 1
				}
				// check DOWN
				if y < len(m)-1 && m[y+1][x] != '#' && nextDirection != 0 {
					directions[1] = 1
				}
				// check LEFT
				if x > 0 && m[y][x-1] != '#' && nextDirection != 3 {
					directions[2] = 1
				}
				// check RIGHT
				if x < len(m[0])-1 && m[y][x+1] != '#' && nextDirection != 2 {
					directions[3] = 1
				}
				sum := 0
				for i := 0; i < len(directions); i++ {
					sum += directions[i]
				}
				if sum == 0 {
					deadQueue[fmt.Sprintf("%d,%d,%d", n.x, n.y, n.direction)] = struct{}{}
					// no way
					break
				} else if sum == 1 {
					// one direction
					for i, d := range directions {
						if d == 1 {
							nextDirection = i
							break
						}
					}
				} else {
					// new node
					for i, d := range directions {
						if d == 1 {
							if _, isDeadWay := deadQueue[fmt.Sprintf("%d,%d,%d", x, y, i)]; !isDeadWay {
								if _, isVisited := n.visited[fmt.Sprintf("%d,%d", x, y)]; !isVisited {
									visited := make(map[string]struct{})
									for key, value := range n.visited {
										visited[key] = value
									}
									visited[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
									queue = append(queue, node2{
										x:         x,
										y:         y,
										step:      step,
										direction: i,
										visited:   visited,
									})
								}
							}
						}
					}
					break
				}
			}
		}
	}
	return maxSteps, nil
}

func main() {
	abs, _ := filepath.Abs("src/day23/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getLongestPathTiles(output, true))
	//fmt.Println(getLongestPathTiles(output, false))
	fmt.Println(getLongestPathTiles2(output))
}
