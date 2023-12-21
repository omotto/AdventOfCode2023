package main

import (
	"advent2023/pkg/math"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"advent2023/pkg/file"
)

func parseInput(i []string) (map[string][]string, error) {
	r := make(map[string][]string)
	for _, l := range i {
		sides := strings.Split(l, " -> ")
		if len(sides) != 2 {
			return nil, errors.New("invalid input")
		}
		r[sides[0]] = strings.Split(sides[1], ", ")
	}
	return r, nil
}

type node struct {
	key string
}

func getPulses(m map[string][]string, numPulses int) (pos int, neg int) {
	var nodes []node
	mem := make(map[string]int)
	for numPulse := 0; numPulse < numPulses; numPulse++ {
		nodes = append(nodes, node{
			key: "broadcaster",
		})
		mem["broadcaster"] = -1
		neg += 1
		for {
			if len(nodes) == 0 {
				break
			} else {
				// pop
				n := nodes[0]
				nodes = nodes[1:]
				//
				var nextNodeKeys []string
				if v1, ok1 := m["&"+n.key]; ok1 {
					nextNodeKeys = v1
				} else if v2, ok2 := m["%"+n.key]; ok2 {
					nextNodeKeys = v2
				} else if n.key == "broadcaster" {
					nextNodeKeys = m[n.key]
				}
				for _, key := range nextNodeKeys {
					if mem[n.key] == 1 {
						pos++
					} else {
						neg++
					}
					if _, exist := m["&"+key]; exist {
						allHigh := true
						for k, v := range m {
							for _, s := range v {
								if s == key {
									nodeValue := -1
									if val, ok := mem[k[1:]]; ok {
										nodeValue = val
									}
									if nodeValue == -1 {
										allHigh = false
										goto out
									}
								}
							}
						}
					out:
						out := 1
						if allHigh {
							out = -1
						}
						mem[key] = out
						nodes = append(nodes, node{
							key: key,
						})
					} else {
						nodeValue := -1
						if v, ok := mem[key]; ok {
							nodeValue = v
						}
						if mem[n.key] == -1 {
							mem[key] = -nodeValue
							nodes = append(nodes, node{
								key: key,
							})
						}
					}
				}
			}
		}
	}
	return pos, neg
}

func getMultipliedPulses(s []string) (int, error) {
	if m, err := parseInput(s); err != nil {
		return 0, err
	} else {
		p, n := getPulses(m, 1000)
		return n * p, nil
	}
}

func getTotalPulses(m map[string][]string) int {
	var nodes []node
	lastStageNodes := []string{"kl", "vm", "kv", "vb"}
	mem := make(map[string]int)
	mainNodesFreq := make(map[string]int)
	numPulse := 0
	for {
		nodes = append(nodes, node{
			key: "broadcaster",
		})
		mem["broadcaster"] = -1
		for {
			if len(nodes) == 0 {
				break
			} else {
				// pop
				n := nodes[0]
				nodes = nodes[1:]
				//
				var nextNodeKeys []string
				if v1, ok1 := m["&"+n.key]; ok1 {
					nextNodeKeys = v1
				} else if v2, ok2 := m["%"+n.key]; ok2 {
					nextNodeKeys = v2
				} else if n.key == "broadcaster" {
					nextNodeKeys = m[n.key]
				}
				for _, key := range nextNodeKeys {
					if _, exist := m["&"+key]; exist {
						allHigh := true
						for k, v := range m {
							for _, s := range v {
								if s == key {
									nodeValue := -1
									if val, ok := mem[k[1:]]; ok {
										nodeValue = val
									}
									if nodeValue == -1 {
										allHigh = false
										goto out
									}
								}
							}
						}
					out:
						out := 1
						if allHigh {
							out = -1
						}
						mem[key] = out
						nodes = append(nodes, node{
							key: key,
						})
						// Added for 2on part calculation
						for _, lastNode := range lastStageNodes {
							if key == lastNode && out == 1 && numPulse > 0 {
								if _, ok := mainNodesFreq[key]; !ok {
									mainNodesFreq[key] = numPulse
								}
							}
						}
						if len(mainNodesFreq) == 4 {
							idx := 0
							freqSteps := make([]int, 4)
							for _, v := range mainNodesFreq {
								freqSteps[idx] = v + 1
								idx++
							}
							return math.LCM(freqSteps[0], freqSteps[1:])
						}
					} else {
						nodeValue := -1
						if v, ok := mem[key]; ok {
							nodeValue = v
						}
						if mem[n.key] == -1 {
							mem[key] = -nodeValue
							nodes = append(nodes, node{
								key: key,
							})
						}
					}
				}
			}
		}
		numPulse++
	}
}

func main() {
	abs, _ := filepath.Abs("src/day20/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getMultipliedPulses(output))
	m, _ := parseInput(output)
	fmt.Println(getTotalPulses(m)) // 241528184647003
}
