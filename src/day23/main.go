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
	maxSteps := 0
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
				if n.step > maxSteps {
					maxSteps = n.step
				}
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

type vector struct {
	sx, sy, dx, dy, steps int
	id                    string
}

func getVectorID(sx, sy, dx, dy, steps int) string {
	if sy < dy {
		return fmt.Sprintf("%d,%d,%d,%d,%d", sx, sy, dx, dy, steps)
	} else if dy < sy {
		return fmt.Sprintf("%d,%d,%d,%d,%d", dx, dy, sx, sy, steps)
	} else if sx < dx {
		return fmt.Sprintf("%d,%d,%d,%d,%d", sx, sy, dx, dy, steps)
	} else if dx < sx {
		return fmt.Sprintf("%d,%d,%d,%d,%d", dx, dy, sx, sy, steps)
	}
	panic("invalid vector")
}

type cross struct {
	x, y, direction int
}

func getVectors(sx, sy, dx, dy int, m []string) []vector {
	var vectors []vector
	visitedNode := map[string]struct{}{fmt.Sprintf("%d,%d", sx, sy): {}}
	queue := []cross{
		{
			x:         sx,
			y:         sy,
			direction: 1,
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
		// search next node
		x := n.x
		y := n.y
		steps := 0
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
			steps++
			// check final node
			if x == dx && y == dy {
				exist := false
				id := getVectorID(n.x, n.y, x, y, steps)
				for _, v := range vectors {
					if v.id == id {
						exist = true
						break
					}
				}
				if !exist {
					vectors = append(vectors, vector{
						sx:    n.x,
						sy:    n.y,
						dx:    x,
						dy:    y,
						steps: steps,
						id:    id,
					})
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
						if _, isVisited := visitedNode[fmt.Sprintf("%d,%d", x, y)]; !isVisited {
							queue = append(queue, cross{
								x:         x,
								y:         y,
								direction: i,
							})
						}
						exist := false
						id := getVectorID(n.x, n.y, x, y, steps)
						for _, v := range vectors {
							if v.id == id {
								exist = true
								break
							}
						}
						if !exist {
							vectors = append(vectors, vector{
								sx:    n.x,
								sy:    n.y,
								dx:    x,
								dy:    y,
								steps: steps,
								id:    id,
							})
						}
					}
				}
				visitedNode[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
				break
			}
		}
	}
	return vectors
}

func getLongestPath(sx, sy, dx, dy int, vectors []vector) int {
	var firstVector vector
	for _, v := range vectors {
		if v.sx == sx && v.sy == sy {
			firstVector = v
			break
		}
	}
	max := 0
	type node struct {
		accSteps int
		dx, dy   int
		visited  map[string]struct{}
	}
	queue := []node{{
		dx:       firstVector.dx,
		dy:       firstVector.dy,
		accSteps: firstVector.steps,
		visited:  map[string]struct{}{firstVector.id: {}},
	}}
	for {
		// exit condition
		if len(queue) == 0 {
			break
		}
		// pop
		n := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// check final node
		if n.dx == dx && n.dy == dy {
			if n.accSteps > max {
				max = n.accSteps
			}
		} else {
			for _, v := range vectors {
				if v.sx == n.dx && v.sy == n.dy || v.dx == n.dx && v.dy == n.dy {
					if _, isVisited := n.visited[v.id]; !isVisited {
						newVisited := make(map[string]struct{})
						for key, val := range n.visited {
							newVisited[key] = val
						}
						newVisited[v.id] = struct{}{}
						var ddx, ddy int
						if v.sx == n.dx && v.sy == n.dy {
							ddx = v.dx
							ddy = v.dy
						} else {
							ddx = v.sx
							ddy = v.sy
						}
						queue = append(queue, node{
							accSteps: v.steps + n.accSteps,
							dx:       ddx,
							dy:       ddy,
							visited:  newVisited,
						})
					}
				}
			}
		}
	}
	return max
}

func getLongestPathTiles3(m []string) (int, error) {
	if sx, sy, e1 := getInitialPoint(m); e1 != nil {
		return -1, e1
	} else if dx, dy, e2 := getFinalPoint(m); e2 != nil {
		return -1, e2
	} else {
		vectors := getVectors(sx, sy, dx, dy, m)
		return getLongestPath(sx, sy, dx, dy, vectors), nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day23/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getLongestPathTiles(output, true))
	//fmt.Println(getLongestPathTiles(output, false))
	fmt.Println(getLongestPathTiles2(output))
	//fmt.Println(getLongestPathTiles3(output))
}
