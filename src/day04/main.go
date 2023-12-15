package main

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type Card struct {
	Id      int
	WinNums []int
	Nums    []int
}

func parseCardStr(s string) (c Card, err error) {
	ss := strings.Join(strings.Fields(s), " ")
	r := strings.Split(ss, ": ")
	if len(r) != 2 {
		return c, errors.New("invalid input string")
	}
	// Get ID
	r0 := strings.Split(r[0], " ")
	if len(r0) != 2 {
		return c, errors.New("invalid input string")
	}
	if strings.ToLower(r0[0]) != "card" {
		return c, errors.New("invalid input string")
	}
	c.Id, err = strconv.Atoi(r0[1])
	if err != nil {
		return c, err
	}
	// Get winner numbers
	r1 := strings.Split(r[1], " | ")
	if len(r1) != 2 {
		return c, errors.New("invalid input string")
	}
	winNums := strings.Split(r1[0], " ")
	var winNum int
	for _, v := range winNums {
		winNum, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
		c.WinNums = append(c.WinNums, winNum)
	}
	// Get numbers
	nums := strings.Split(r1[1], " ")
	var num int
	for _, v := range nums {
		num, err = strconv.Atoi(v)
		if err != nil {
			return c, err
		}
		c.Nums = append(c.Nums, num)
	}
	return c, err
}

func getNumWinnerNumbers(c Card) (winNums []int) {
	for w := 0; w < len(c.WinNums); w++ {
		for n := 0; n < len(c.Nums); n++ {
			if c.Nums[n] == c.WinNums[w] {
				winNums = append(winNums, c.Nums[n])
			}
		}
	}
	return winNums
}

func getPowFromCard(winNums []int) int {
	return int(math.Pow(2, float64(len(winNums)-1)))
}

func getSumOfAllCards(s []string) (sum int, err error) {
	var c Card
	for _, l := range s {
		if c, err = parseCardStr(l); err != nil {
			return sum, err
		} else {
			winNums := getNumWinnerNumbers(c)
			sum += getPowFromCard(winNums)
		}
	}
	return sum, err
}

func getNumWinnerCards(s []string, l int, a int) (num int) {
	card, _ := parseCardStr(s[l])
	numWinners := len(getNumWinnerNumbers(card))
	if numWinners == 0 {
		return 0
	} else {
		var acc int = 0
		for i := 0; i < numWinners; i++ {
			acc += getNumWinnerCards(s, l+i+1, a+1)
		}
		return numWinners + acc
	}
}

func getTotalWinnerCards(s []string) (num int) {
	for i, _ := range s {
		num += 1 + getNumWinnerCards(s, i, 0)
	}
	return num
}

func main() {
	abs, _ := filepath.Abs("src/day04/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumOfAllCards(output))
	fmt.Println(getTotalWinnerCards(output))
}
