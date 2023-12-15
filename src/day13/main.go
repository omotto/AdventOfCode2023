package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
)

func getPattern(s []string, index int) (p []string, n int) {
	var i = index
	for {
		if i < len(s)-1 {
			if len(s[i]) > 1 {
				p = append(p, s[i])
			} else {
				break
			}
		} else {
			break
		}
		i += 1
	}
	return p, i + 1
}

func getMinDistance(s []int, d int) (r int) {
	var shortest int = 9999999999
	for _, source := range s {
		distance := source - d
		if distance < 0 {
			distance *= -1
		}
		if distance < shortest {
			shortest = distance
			r = source
		}
	}
	return r
}

func getMirrorNodes(p []string) (v, h int) {
	var (
		matches []int
		match   bool
	)
	// vertical
	for i := 0; i < len(p)-1; i++ {
		match = true
		for j := i; j >= 0; j-- {
			if i+i+1-j > len(p)-1 {
				break
			}
			if p[j] != p[i+i+1-j] {
				match = false
				break
			}
		}
		if match {
			matches = append(matches, i+1)
		}
	}
	h = getMinDistance(matches, len(p)/2)
	//	horizontal
	matches = []int{}
	for i := 0; i < len(p[0])-1; i++ {
		match = true
		for j := i; j >= 0; j-- {
			if i+i+1-j > len(p[0])-1 {
				break
			}
			for y := 0; y < len(p); y++ {
				if p[y][j] != p[y][i+i+1-j] {
					match = false
					break
				}
			}
		}
		if match {
			matches = append(matches, i+1)
		}
	}
	v = getMinDistance(matches, len(p)/2)
	return v, h
}

func getMirrorNodes2(p []string) (v, h int) {
	//fv, fh := getMirrorNodes(p)
	var (
		listChanges []int
		numChanges  int
	)
	for i := 0; i < len(p[0])-1; i++ {
		numChanges = 0
		for j := i; j >= 0; j-- {
			if i+i+1-j > len(p[0])-1 {
				break
			}
			for y := 0; y < len(p); y++ {
				if p[y][j] != p[y][i+i+1-j] {
					numChanges++
				}
			}
		}
		if numChanges == 1 {
			listChanges = append(listChanges, i+1)
		}
	}
	v = getMinDistance(listChanges, len(p)/2)
	listChanges = []int{}
	for i := 0; i < len(p)-1; i++ {
		numChanges = 0
		for j := i; j >= 0; j-- {
			if i+i+1-j > len(p)-1 {
				break
			}
			for x := 0; x < len(p[0]); x++ {
				if p[j][x] != p[i+i+1-j][x] {
					numChanges++
				}
			}
		}
		if numChanges == 1 {
			listChanges = append(listChanges, i+1)
		}
	}
	h = getMinDistance(listChanges, len(p)/2)

	return v, h
}

func getSumNotes(s []string, getMirrorNodes func(p []string) (v, h int)) (sum int) {
	var (
		next    int = 0
		pattern []string
		v, h    int
	)
	for {
		pattern, next = getPattern(s, next)
		v, h = getMirrorNodes(pattern)
		sum += h*100 + v
		if next > len(s)-1 {
			break
		}
	}
	return sum
}

func main() {
	abs, _ := filepath.Abs("src/day13/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumNotes(output, getMirrorNodes))
	fmt.Println(getSumNotes(output, getMirrorNodes2))
}
