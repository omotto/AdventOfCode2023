package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
	"strconv"
)

func getValidNumbers(s []string) (nums []int) {
	var (
		firstPos   int = -1
		lastPos    int = -1
		currentNum int = 0
	)
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] >= '0' && s[y][x] <= '9' {
				num, _ := strconv.Atoi(fmt.Sprintf("%c", s[y][x]))
				currentNum = currentNum*10 + num
				if firstPos == -1 {
					firstPos = x
				}
				lastPos = x
			} else {
				if firstPos != -1 {
					valid := false
					// left
					if firstPos-1 >= 0 {
						if s[y][firstPos-1] != '.' {
							valid = true
						}
					}
					// right
					if lastPos+1 < len(s[y]) {
						if s[y][lastPos+1] != '.' {
							valid = true
						}
					}
					// up
					if y-1 >= 0 {
						for i := firstPos; i < lastPos+1; i++ {
							if s[y-1][i] != '.' {
								valid = true
							}
						}
					}
					// down
					if y+1 < len(s) {
						for i := firstPos; i < lastPos+1; i++ {
							if s[y+1][i] != '.' {
								valid = true
							}
						}
					}
					// upper left
					if firstPos-1 >= 0 && y-1 >= 0 {
						if s[y-1][firstPos-1] != '.' {
							valid = true
						}
					}
					// upper right
					if lastPos+1 < len(s[y]) && y-1 >= 0 {
						if s[y-1][lastPos+1] != '.' {
							valid = true
						}
					}
					// down left
					if firstPos-1 >= 0 && y+1 < len(s) {
						if s[y+1][firstPos-1] != '.' {
							valid = true
						}
					}
					// down right
					if lastPos+1 < len(s[y]) && y+1 < len(s) {
						if s[y+1][lastPos+1] != '.' {
							valid = true
						}
					}
					//
					if valid {
						nums = append(nums, currentNum)
					}
				}
				firstPos = -1
				currentNum = 0
			}
		}
	}
	return nums
}

func getGearNumbers(s []string) (nums []int) {
	type number struct {
		numVal int
		numX   int
		numY   int
		symX   int
		symY   int
	}
	var (
		firstPos   int = -1
		lastPos    int = -1
		currentNum int = 0
		numbers    []number
	)
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] >= '0' && s[y][x] <= '9' {
				num, _ := strconv.Atoi(fmt.Sprintf("%c", s[y][x]))
				currentNum = currentNum*10 + num
				if firstPos == -1 {
					firstPos = x
				}
				lastPos = x
			} else {
				if firstPos != -1 {
					// left
					if firstPos-1 >= 0 {
						if s[y][firstPos-1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y,
								symX:   firstPos - 1,
							})
						}
					}
					// right
					if lastPos+1 < len(s[y]) {
						if s[y][lastPos+1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y,
								symX:   lastPos + 1,
							})
						}
					}
					// up
					if y-1 >= 0 {
						for i := firstPos; i < lastPos+1; i++ {
							if s[y-1][i] == '*' {
								numbers = append(numbers, number{
									numVal: currentNum,
									numX:   firstPos,
									numY:   y,
									symY:   y - 1,
									symX:   i,
								})
							}
						}
					}
					// down
					if y+1 < len(s) {
						for i := firstPos; i < lastPos+1; i++ {
							if s[y+1][i] == '*' {
								numbers = append(numbers, number{
									numVal: currentNum,
									numX:   firstPos,
									numY:   y,
									symY:   y + 1,
									symX:   i,
								})
							}
						}
					}
					// upper left
					if firstPos-1 >= 0 && y-1 >= 0 {
						if s[y-1][firstPos-1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y - 1,
								symX:   firstPos - 1,
							})
						}
					}
					// upper right
					if lastPos+1 < len(s[y]) && y-1 >= 0 {
						if s[y-1][lastPos+1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y - 1,
								symX:   lastPos + 1,
							})
						}
					}
					// down left
					if firstPos-1 >= 0 && y+1 < len(s) {
						if s[y+1][firstPos-1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y + 1,
								symX:   firstPos - 1,
							})
						}
					}
					// down right
					if lastPos+1 < len(s[y]) && y+1 < len(s) {
						if s[y+1][lastPos+1] == '*' {
							numbers = append(numbers, number{
								numVal: currentNum,
								numX:   firstPos,
								numY:   y,
								symY:   y + 1,
								symX:   lastPos + 1,
							})
						}
					}
				}
				firstPos = -1
				currentNum = 0
			}
		}
	}
	// Check numbers
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j != i &&
				numbers[i].symY == numbers[j].symY &&
				numbers[i].symX == numbers[j].symX &&
				(numbers[i].numX != numbers[j].numX || numbers[i].numY != numbers[j].numY) {
				nums = append(nums, numbers[i].numVal*numbers[j].numVal)
				// disable j value
				numbers[j].numY = numbers[i].numY
				numbers[j].numX = numbers[i].numX
			}
		}
	}
	return nums
}

func getSumValidNumsFromMap(m []string) (sum int) {
	validNumbers := getValidNumbers(m)
	for _, num := range validNumbers {
		sum += num
	}
	return sum
}

func getSumGearNumsFromMap(m []string) (sum int) {
	validNumbers := getGearNumbers(m)
	for _, num := range validNumbers {
		sum += num
	}
	return sum
}

func main() {
	abs, _ := filepath.Abs("src/day03/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumValidNumsFromMap(output))
	fmt.Println(getSumGearNumsFromMap(output))
}
