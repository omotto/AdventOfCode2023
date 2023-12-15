package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type raceInfo struct {
	times     []int
	distances []int
}

func parseInput(s []string) (r raceInfo, err error) {
	if len(s) != 2 {
		return r, errors.New("invalid input")
	}
	s0 := strings.Join(strings.Fields(s[0]), " ")
	s1 := strings.Join(strings.Fields(s[1]), " ")
	ss0 := strings.Split(s0, ": ")
	if len(ss0) != 2 {
		return r, errors.New("invalid input")
	}
	if strings.ToLower(ss0[0]) != "time" {
		return r, errors.New("invalid input")
	}
	times := strings.Split(ss0[1], " ")
	r.times = make([]int, len(times))
	for i, time := range times {
		if r.times[i], err = strconv.Atoi(time); err != nil {
			return r, errors.New("invalid input")
		}
	}
	ss1 := strings.Split(s1, ": ")
	if len(ss1) != 2 {
		return r, errors.New("invalid input")
	}
	if strings.ToLower(ss1[0]) != "distance" {
		return r, errors.New("invalid input")
	}
	distances := strings.Split(ss1[1], " ")
	r.distances = make([]int, len(distances))
	for i, distance := range distances {
		if r.distances[i], err = strconv.Atoi(distance); err != nil {
			return r, errors.New("invalid input")
		}
	}
	return r, err
}

func calculateDistance(time, maxTime int) (distance int) {
	distance = (maxTime - time) * time
	return distance
}

func calculateWinners(maxTime, duration int) (winners []int) {
	var isWinner bool = false
	if maxTime > 1 {
		for i := 1; i < maxTime; i++ {
			if calculateDistance(i, maxTime) > duration {
				winners = append(winners, i)
				isWinner = true
			} else if isWinner {
				break
			}
		}
	}
	return winners
}

func getResults(s []string) (r int, err error) {
	var info raceInfo
	if info, err = parseInput(s); err != nil {
		return r, err
	} else {
		r = 1
		for idx, time := range info.times {
			r *= len(calculateWinners(time, info.distances[idx]))
		}
	}
	return r, err
}

func main() {
	abs, _ := filepath.Abs("src/day06/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getResults(output))
	fmt.Println(len(calculateWinners(61709066, 643118413621041)))
}
