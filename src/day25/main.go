package main

import (
	"advent2023/pkg/file"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func parseInput(in []string) (map[string][]string, error) {
	nodes := make(map[string][]string)
	for _, line := range in {
		s := strings.Split(line, ": ")
		if len(s) != 2 {
			return nil, errors.New("invalid input")
		}
		src := s[0]
		dst := strings.Split(s[1], " ")
		nodes[src] = append(nodes[src], dst...)
		for _, d := range dst {
			nodes[d] = append(nodes[d], src)
		}
	}
	return nodes, nil
}

// Swept all nodes straight from every node to get how many times one vector is used
func getWeightedVectors(nodes map[string][]string) map[string]int {
	vectors := map[string]int{}
	for node, _ := range nodes {
		visited := map[string]struct{}{}
		queue := []string{node}
		for {
			if len(queue) == 0 {
				break
			}
			nextNode := queue[0]
			queue = queue[1:]
			for _, dstNode := range nodes[nextNode] {
				if _, isVisited := visited[dstNode]; !isVisited {
					visited[dstNode] = struct{}{}
					queue = append(queue, dstNode)
					sortKey := []string{nextNode, dstNode}
					sort.Strings(sortKey)
					key := fmt.Sprintf("%s,%s", sortKey[0], sortKey[1])
					vectors[key]++
				}
			}
		}
	}
	return vectors
}

func removeVectorFromNodes(nodes map[string][]string, maxVector string) {
	// Remove from nodes
	s := strings.Split(maxVector, ",")
	src := s[0]
	dst := s[1]
	var newDst, newSrc []string
	//
	for _, v := range nodes[src] {
		if v != dst {
			newDst = append(newDst, v)
		}
	}
	nodes[src] = newDst
	//
	for _, v := range nodes[dst] {
		if v != src {
			newSrc = append(newSrc, v)
		}
	}
	nodes[dst] = newSrc
}

func countNodes(nodes map[string][]string) int {
	for node, _ := range nodes {
		visited := map[string]struct{}{}
		queue := []string{node}
		for {
			if len(queue) == 0 {
				break
			}
			nextNode := queue[0]
			queue = queue[1:]
			for _, dstNode := range nodes[nextNode] {
				if _, isVisited := visited[dstNode]; !isVisited {
					visited[dstNode] = struct{}{}
					queue = append(queue, dstNode)
				}
			}
		}
		return len(visited)
	}
	return 0
}

func getMultipliedSize(in []string) (int, error) {
	if nodes, err := parseInput(in); err != nil {
		return 0, err
	} else {
		var maxVector string
		totalNodes := len(nodes)
		for idx := 0; idx < 3; idx++ {
			vectors := getWeightedVectors(nodes)
			// Get max swept vector
			max := 0
			for k, v := range vectors {
				if v > max {
					max = v
					maxVector = k
				}
			}
			// remove it
			removeVectorFromNodes(nodes, maxVector)
		}
		firstGroup := countNodes(nodes) // fewer nodes due to not connected
		secondGroup := totalNodes - firstGroup
		return firstGroup * secondGroup, nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day25/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getMultipliedSize(output))
}
