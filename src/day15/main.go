package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

func parseInputLine(s string) []string {
	return strings.Split(s, ",")
}

func calculateHash(s string) int {
	acc := 0
	for i := 0; i < len(s); i++ {
		acc = ((int(s[i]) + acc) * 17) % 256
	}
	return acc
}

func getSumHash(s []string) int {
	sum := 0
	cache := make(map[string]int)
	values := parseInputLine(s[0])
	for _, v := range values {
		if n, ok := cache[v]; ok {
			sum += n
		} else {
			n = calculateHash(v)
			sum += n
			cache[v] = n
		}
	}
	return sum
}

type keyValue struct {
	label string
	focus int
}

func getFocusingPower(s []string) int {
	sum := 0
	cache := make(map[string]int)
	boxes := make([][]keyValue, 256)
	values := parseInputLine(s[0])
	for _, v := range values {
		if v[len(v)-1] == '-' {
			label := v[:len(v)-1]
			box, ok := cache[label]
			if !ok {
				box = calculateHash(label)
			}
			for idx, boxValue := range boxes[box] {
				if boxValue.label == label {
					boxes[box][idx].focus = -1
					boxes[box][idx].label = ""
					break
				}
			}
		} else {
			r := strings.Split(v, "=")
			label := r[0]
			box, ok := cache[label]
			if !ok {
				box = calculateHash(label)
			}
			focus, _ := strconv.Atoi(r[1])
			found := false
			for idx, boxValue := range boxes[box] {
				if boxValue.label == label {
					boxes[box][idx].focus = focus
					found = true
					break
				}
			}
			if !found {
				boxes[box] = append(boxes[box], keyValue{
					label: label,
					focus: focus,
				})
			}
		}
	}
	for i, box := range boxes {
		pos := 0
		for _, value := range box {
			if value.label != "" {
				sum += (i + 1) * (pos + 1) * value.focus
				pos++
			}
		}
	}
	return sum
}

func main() {
	abs, _ := filepath.Abs("src/day15/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumHash(output))
	fmt.Println(getFocusingPower(output))
}
