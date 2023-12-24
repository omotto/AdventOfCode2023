package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2023/pkg/file"
)

type hailstone struct {
	x, y, z, dx, dy, dz float64
}

func parseEquations(in []string) ([]hailstone, error) {
	//280342553482089, 252447176665186, 220895081604913 @ 127, 12, 96
	var hailstones []hailstone
	for _, line := range in {
		parsedLine := strings.TrimSpace(strings.Replace(line, " ", "", -1))
		a := strings.Split(parsedLine, "@")
		if len(a) != 2 {
			return nil, errors.New("invalid input")
		}
		coordinates := strings.Split(a[0], ",")
		velocity := strings.Split(a[1], ",")
		if len(coordinates) != 3 || len(velocity) != 3 {
			return nil, errors.New("invalid input")
		}
		x, ex := strconv.Atoi(coordinates[0])
		y, ey := strconv.Atoi(coordinates[1])
		z, ez := strconv.Atoi(coordinates[2])
		if ex != nil || ey != nil || ez != nil {
			return nil, errors.New("invalid input")
		}
		dx, edx := strconv.Atoi(velocity[0])
		dy, edy := strconv.Atoi(velocity[1])
		dz, edz := strconv.Atoi(velocity[2])
		if edx != nil || edy != nil || edz != nil {
			return nil, errors.New("invalid input")
		}
		hailstones = append(hailstones, hailstone{
			x:  float64(x),
			y:  float64(y),
			z:  float64(z),
			dx: float64(dx),
			dy: float64(dy),
			dz: float64(dz),
		})
	}
	return hailstones, nil
}

func getSlope(h hailstone) float64 {
	return h.dy / h.dx
}

func areParallel(h1, h2 hailstone) bool {
	return getSlope(h1) == getSlope(h2)
}

func areCrossed(h1, h2 hailstone) bool {
	cx := (h2.y - (h2.dy/h2.dx)*h2.x - (h1.y - (h1.dy/h1.dx)*h1.x)) / (h1.dy/h1.dx - h2.dy/h2.dx)
	return !((cx > h1.x) == (h1.dx > 0) && (cx > h2.x) == (h2.dx > 0))
}

func calculateEquation(h1, h2 hailstone) (float64, float64) {
	if areParallel(h1, h2) {
		return -1, -1
	}
	// Line Equation X*dy - Y*dx = x*dy -y*dx
	if areCrossed(h1, h2) {
		return -1, -1
	}
	// Solve equation
	y := ((-h1.dy/h2.dy)*(h2.x*h2.dy-h2.y*h2.dx) + h1.x*h1.dy - h1.y*h1.dx) / ((h2.dx * h1.dy / h2.dy) - h1.dx)
	x := y*h2.dx/h2.dy + (h2.x*h2.dy-h2.y*h2.dx)/h2.dy
	return x, y
}

func getCrossedHailstones(in []string, min, max float64) (int, error) {
	if hailstones, err := parseEquations(in); err != nil {
		return -1, err
	} else {
		count := 0
		for i := 0; i < len(hailstones)-1; i++ {
			for j := i + 1; j < len(hailstones); j++ {
				x, y := calculateEquation(hailstones[i], hailstones[j])
				if x >= min && x <= max && y >= min && y <= max {
					count++
				}
			}
		}
		return count, nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day24/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getCrossedHailstones(output, 200000000000000.0, 400000000000000.0))
}
