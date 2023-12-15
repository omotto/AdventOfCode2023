package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

func checkDigit(s string, firstLast bool, numOnly bool) (n int, err error) {
	if firstLast {
		if s[0] >= '0' && s[0] <= '9' {
			return strconv.Atoi(fmt.Sprintf("%c", s[0]))
		}
	} else {
		if s[len(s)-1] >= '0' && s[len(s)-1] <= '9' {
			return strconv.Atoi(fmt.Sprintf("%c", s[len(s)-1]))
		}
	}
	if !numOnly {
		stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		if firstLast {
			for idx, stringDigit := range stringDigits {
				if strings.HasPrefix(strings.ToLower(s), stringDigit) {
					return idx + 1, err
				}
			}
		} else {
			for idx, stringDigit := range stringDigits {
				if strings.HasSuffix(strings.ToLower(s), stringDigit) {
					return idx + 1, err
				}
			}
		}
	}
	return -1, err
}

func getFirsLastDigits(s string, numOnly bool) (n int, err error) {
	var (
		first  int = -1
		last   int = -1
		lenStr int = len(s)
	)
	for i := 0; i < lenStr; i++ {
		if first == -1 {
			if first, err = checkDigit(s[i:], true, numOnly); err != nil {
				return 0, err
			}
		}
		if last == -1 {
			if last, err = checkDigit(s[:lenStr-i], false, numOnly); err != nil {
				return 0, err
			}
		}
		if last > -1 && first > -1 {
			return first*10 + last, err
		}
	}
	return 0, err
}

func getSumCalibrationNumbers(l []string, numOnly bool) (n int, err error) {
	var v int = 0
	for _, s := range l {
		if v, err = getFirsLastDigits(s, numOnly); err != nil {
			return 0, err
		}
		n += v
	}
	return n, err
}

func main() {
	abs, _ := filepath.Abs("src/day01/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSumCalibrationNumbers(output, true))
	fmt.Println(getSumCalibrationNumbers(output, false))
}
