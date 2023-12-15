package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

func parseInput(s []string) (l [][]int, err error) {
	length := len(s)
	l = make([][]int, length)
	for i, vl := range s {
		r := strings.Split(vl, " ")
		l[i] = make([]int, len(r))
		for j, vn := range r {
			if l[i][j], err = strconv.Atoi(vn); err != nil {
				return l, err
			}
		}
	}
	return l, err
}

func calculatePrevNumber(s []int) (r int) {
	var finish bool = true
	ss := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		ss[i] = s[i+1] - s[i]
		if ss[i] != 0 {
			finish = false
		}
	}
	if finish {
		return s[0]
	}
	return s[0] - calculatePrevNumber(ss)
}

func calculateNextNumber(s []int) (r int) {
	var finish bool = true
	ss := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		ss[i] = s[i+1] - s[i]
		if ss[i] != 0 {
			finish = false
		}
	}
	if finish {
		return s[0]
	}
	return s[len(s)-1] + calculateNextNumber(ss)
}

func getSumOASIS(s []string) (result int, err error) {
	var listOfSequences [][]int
	if listOfSequences, err = parseInput(s); err != nil {
		return result, err
	}
	for _, sequence := range listOfSequences {
		result += calculateNextNumber(sequence)
	}
	return result, err
}

func getSumOASIS2(s []string) (result int, err error) {
	var listOfSequences [][]int
	if listOfSequences, err = parseInput(s); err != nil {
		return result, err
	}
	for _, sequence := range listOfSequences {
		result += calculatePrevNumber(sequence)
	}
	return result, err
}

func main() {
	abs, _ := filepath.Abs("src/day09/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumOASIS(output))
	fmt.Println(getSumOASIS2(output))
}
