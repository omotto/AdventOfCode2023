package main

import (
	"advent2023/pkg/file"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type unitConverse struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type converseMap struct {
	seedToSoil            []unitConverse
	soilToFertilizer      []unitConverse
	fertilizerToWater     []unitConverse
	waterToLight          []unitConverse
	lightToTemperature    []unitConverse
	temperatureToHumidity []unitConverse
	humidityToLocation    []unitConverse
}

func getSeedsPart1(s string) (seeds []int, err error) {
	r := strings.Split(s, ": ")
	if len(r) != 2 {
		return nil, errors.New("invalid input")
	}
	ss := strings.Split(r[1], " ")
	var seed int
	for _, v := range ss {
		seed, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, seed)
	}
	return seeds, err
}

func getSeedsPart2(s string) (seeds []int, err error) {
	r := strings.Split(s, ": ")
	if len(r) != 2 {
		return nil, errors.New("invalid input")
	}
	ss := strings.Split(r[1], " ")
	if len(ss)%2 != 0 {
		return nil, errors.New("invalid input")
	}
	var start, length int
	for i := 0; i < len(ss); i += 2 {
		start, err = strconv.Atoi(ss[i])
		if err != nil {
			return nil, err
		}
		length, err = strconv.Atoi(ss[i+1])
		if err != nil {
			return nil, err
		}
		for j := start; j < start+length; j++ {
			seeds = append(seeds, j)
		}
	}
	return seeds, err
}

func getValues(s string) (value unitConverse, err error) {
	r := strings.Split(s, " ")
	if len(r) != 3 {
		return value, errors.New("invalid input")
	}
	value.destinationRangeStart, err = strconv.Atoi(r[0])
	if err != nil {
		return value, err
	}
	value.sourceRangeStart, err = strconv.Atoi(r[1])
	if err != nil {
		return value, err
	}
	value.rangeLength, err = strconv.Atoi(r[2])
	return value, err
}

func parseInput(s []string, getSeeds func(s string) (seeds []int, err error)) (seeds []int, c converseMap, err error) {
	var (
		block  int = 0
		values unitConverse
	)
	for i, v := range s {
		if i == 0 {
			if seeds, err = getSeeds(v); err != nil {
				return nil, c, err
			}
		} else {
			if len(v) > 2 { // discard blank lines
				switch {
				case strings.HasPrefix(v, "seed-to-soil map:"):
					block = 1
				case strings.HasPrefix(v, "soil-to-fertilizer map:"):
					block = 2
				case strings.HasPrefix(v, "fertilizer-to-water map:"):
					block = 3
				case strings.HasPrefix(v, "water-to-light map:"):
					block = 4
				case strings.HasPrefix(v, "light-to-temperature map:"):
					block = 5
				case strings.HasPrefix(v, "temperature-to-humidity map:"):
					block = 6
				case strings.HasPrefix(v, "humidity-to-location map:"):
					block = 7
				default:
					if values, err = getValues(v); err != nil {
						return nil, c, err
					}
					switch block {
					case 1:
						c.seedToSoil = append(c.seedToSoil, values)
					case 2:
						c.soilToFertilizer = append(c.soilToFertilizer, values)
					case 3:
						c.fertilizerToWater = append(c.fertilizerToWater, values)
					case 4:
						c.waterToLight = append(c.waterToLight, values)
					case 5:
						c.lightToTemperature = append(c.lightToTemperature, values)
					case 6:
						c.temperatureToHumidity = append(c.temperatureToHumidity, values)
					case 7:
						c.humidityToLocation = append(c.humidityToLocation, values)
					}
				}
			}
		}
	}
	return seeds, c, err
}

func converse(val int, mc []unitConverse) (ret int) {
	for _, m := range mc {
		if val >= m.sourceRangeStart && val <= m.sourceRangeStart+m.rangeLength {
			return m.destinationRangeStart + val - m.sourceRangeStart
		}
	}
	return val
}

func calculateLocation(seed int, m converseMap) (location int) {
	soil := converse(seed, m.seedToSoil)
	fertilizer := converse(soil, m.soilToFertilizer)
	water := converse(fertilizer, m.fertilizerToWater)
	light := converse(water, m.waterToLight)
	temperature := converse(light, m.lightToTemperature)
	humidity := converse(temperature, m.temperatureToHumidity)
	location = converse(humidity, m.humidityToLocation)
	return location
}

func getMinimumLocation(locations []int) (min int) {
	for i, loc := range locations {
		if i == 0 {
			min = loc
		} else {
			if loc < min {
				min = loc
			}
		}
	}
	return min
}

func getLowerLocationNumber(s []string, getSeeds func(s string) (seeds []int, err error)) (int, error) {
	seeds, m, err := parseInput(s, getSeeds)
	if err != nil {
		return 0, err
	}
	locs := make([]int, len(seeds))
	for i, seed := range seeds {
		locs[i] = calculateLocation(seed, m)
	}
	return getMinimumLocation(locs), nil
}

func main() {
	abs, _ := filepath.Abs("src/day05/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getLowerLocationNumber(output, getSeedsPart1))
	fmt.Println(getLowerLocationNumber(output, getSeedsPart2)) // 50716416 ¿?
}
