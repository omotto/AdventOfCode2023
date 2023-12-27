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
		for _, d := range dst {
			nodes[d] = append(nodes[d], src)
			nodes[src] = append(nodes[src], d)
		}
	}
	return nodes, nil
}

// Swept all nodes from every node to get how many times one vector is used
func sweptAlVectors(nodes map[string][]string) map[string]int {
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

func removeVectorWithMoreSweeps(nodes map[string][]string, vectors map[string]int) string {
	var (
		maxVector string
		max       int = 0
	)
	// Get max swept vector
	for k, v := range vectors {
		if v > max {
			max = v
			maxVector = k
		}
	}
	// Remove from vectors
	delete(vectors, maxVector)
	// Remove from nodes
	s := strings.Split(maxVector, ",")
	src := s[0]
	dst := s[1]
	var newDst, newSrc []string
	for _, v := range nodes[src] {
		if v != dst {
			newDst = append(newDst, v)
		}
	}
	nodes[src] = newDst
	for _, v := range nodes[dst] {
		if v != src {
			newSrc = append(newSrc, v)
		}
	}
	nodes[dst] = newSrc
	return maxVector
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
		totalNodes := len(nodes)
		for idx := 0; idx < 3; idx++ {
			vectors := sweptAlVectors(nodes)
			removeVectorWithMoreSweeps(nodes, vectors)
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
