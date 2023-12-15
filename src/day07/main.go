package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type hand struct {
	cards string
	bid   int
}

func parseInput(s []string) (hands []hand, err error) {
	var bid int
	for _, v := range s {
		sv := strings.Split(v, " ")
		if len(sv) != 2 {
			return hands, errors.New("invalid input")
		}
		if bid, err = strconv.Atoi(sv[1]); err != nil {
			return hands, errors.New("invalid input")
		}
		hands = append(hands, hand{
			cards: sv[0],
			bid:   bid,
		})
	}
	return hands, err
}

func getHandValuePart1(s string) (v int) {
	type rep struct {
		cardType string
		numTimes int
	}
	var (
		repeated []rep
		found    bool
	)
	for i := 0; i < len(s); i++ {
		found = false
		for j := 0; j < len(repeated); j++ {
			if repeated[j].cardType == string(s[i]) {
				found = true
				break
			}
		}
		if !found {
			repeated = append(repeated, rep{
				cardType: string(s[i]),
				numTimes: strings.Count(s, string(s[i])),
			})
		}
	}
	// Five of a kind
	if len(repeated) == 1 {
		return 7
	}
	if len(repeated) == 2 {
		// Four of a kind
		if repeated[0].numTimes == 4 || repeated[1].numTimes == 4 {
			return 6
		}
		// Full house
		return 5
	}
	if len(repeated) == 3 {
		// Four of a kind
		if repeated[0].numTimes == 3 || repeated[1].numTimes == 3 || repeated[2].numTimes == 3 {
			return 4
		}
		// Two pair
		return 3
	}
	// One pair
	if len(repeated) == 4 {
		return 2
	}
	// High card
	return 1
}

func getHandValuePart2(s string) (v int) {
	type rep struct {
		cardType string
		numTimes int
	}
	var (
		repeated []rep
		found    bool
	)
	for i := 0; i < len(s); i++ {
		found = false
		for j := 0; j < len(repeated); j++ {
			if repeated[j].cardType == string(s[i]) {
				found = true
				break
			}
		}
		if !found {
			repeated = append(repeated, rep{
				cardType: string(s[i]),
				numTimes: strings.Count(s, string(s[i])),
			})
		}
	}
	// Five of a kind
	if len(repeated) == 1 {
		return 7
	}
	// Check Jokes
	var (
		maxRepeatedTimes int = 0
		mostRepeatedCard int
		jockerIndex      int
		jockerTimes      int = 0
	)
	for idx, r := range repeated {
		if r.numTimes > maxRepeatedTimes && strings.ToUpper(r.cardType) != "J" {
			maxRepeatedTimes = r.numTimes
			mostRepeatedCard = idx
		}
	}
	for idx, r := range repeated {
		if strings.ToUpper(r.cardType) == "J" {
			jockerIndex = idx
			jockerTimes = r.numTimes
			break
		}
	}
	if jockerTimes > 0 {
		repeated[mostRepeatedCard].numTimes += jockerTimes
		repeated = append(repeated[:jockerIndex], repeated[jockerIndex+1:]...)
	}
	// Five of a kind
	if len(repeated) == 1 {
		return 7
	}
	if len(repeated) == 2 {
		// Four of a kind
		if repeated[0].numTimes == 4 || repeated[1].numTimes == 4 {
			return 6
		}
		// Full house
		return 5
	}
	if len(repeated) == 3 {
		// Four of a kind
		if repeated[0].numTimes == 3 || repeated[1].numTimes == 3 || repeated[2].numTimes == 3 {
			return 4
		}
		// Two pair
		return 3
	}
	// One pair
	if len(repeated) == 4 {
		return 2
	}
	// High card
	return 1
}

func getCardValuePart1(s string) (v int) {
	switch strings.ToUpper(s) {
	case "A":
		return 13
	case "K":
		return 12
	case "Q":
		return 11
	case "J":
		return 10
	case "T":
		return 9
	case "9":
		return 8
	case "8":
		return 7
	case "7":
		return 6
	case "6":
		return 5
	case "5":
		return 4
	case "4":
		return 3
	case "3":
		return 2
	case "2":
		return 1
	default:
		return 0
	}
}

func getCardValuePart2(s string) (v int) {
	switch strings.ToUpper(s) {
	case "A":
		return 13
	case "K":
		return 12
	case "Q":
		return 11
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	case "J":
		return 1
	default:
		return 0
	}
}

func isLess(hand1, hand2 string, getCardValue, getHandValue func(s string) (v int)) bool {
	h1 := getHandValue(hand1)
	h2 := getHandValue(hand2)
	if h1 != h2 {
		return h1 < h2
	} else {
		for i := 0; i < len(hand1); i++ {
			if getCardValue(string(hand1[i])) != getCardValue(string(hand2[i])) {
				return getCardValue(string(hand1[i])) < getCardValue(string(hand2[i]))
			}
		}
	}
	return false
}

func getRankedHands(hands []hand, getCardValue, getHandValue func(s string) (v int)) []hand {
	sort.SliceStable(hands, func(i, j int) bool {
		return isLess(hands[i].cards, hands[j].cards, getCardValue, getHandValue)
	})
	return hands
}

func getTotalWinnings(s []string, getCardValue, getHandValue func(s string) (v int)) (result int, err error) {
	var hands []hand
	if hands, err = parseInput(s); err != nil {
		return result, err
	}
	handsRanked := getRankedHands(hands, getCardValue, getHandValue)
	for i, r := range handsRanked {
		result += r.bid * (i + 1)
	}
	return result, err
}

func main() {
	abs, _ := filepath.Abs("src/day07/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getTotalWinnings(output, getCardValuePart1, getHandValuePart1))
	fmt.Println(getTotalWinnings(output, getCardValuePart2, getHandValuePart2))
}
