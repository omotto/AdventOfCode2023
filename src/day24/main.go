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

func intersectXY(h1, h2 hailstone, vx, vy float64) ([2]float64, error) {
	x1 := h1.x
	y1 := h1.y
	x2 := h1.x + h1.dx - vx
	y2 := h1.y + h1.dy - vy
	x3 := h2.x
	y3 := h2.y
	x4 := h2.x + h2.dx - vx
	y4 := h2.y + h2.dy - vy
	if denominator := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1); denominator == 0 {
		return [2]float64{}, errors.New("invalid denominator")
	} else {
		ua := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denominator
		return [2]float64{x1 + ua*(x2-x1), y1 + ua*(y2-y1)}, nil
	}
}

func intersectYZ(h1, h2 hailstone, vy, vz float64) ([2]float64, error) {
	y1 := h1.y
	z1 := h1.z
	y2 := h1.y + h1.dy - vy
	z2 := h1.z + h1.dz - vz
	y3 := h2.y
	z3 := h2.z
	y4 := h2.y + h2.dy - vy
	z4 := h2.z + h2.dz - vz
	if denominator := (z4-z3)*(y2-y1) - (y4-y3)*(z2-z1); denominator == 0 {
		return [2]float64{}, errors.New("invalid denominator")
	} else {
		ua := ((y4-y3)*(z1-z3) - (z4-z3)*(y1-y3)) / denominator
		return [2]float64{y1 + ua*(y2-y1), z1 + ua*(z2-z1)}, nil
	}
}

func intersectXZ(h1, h2 hailstone, vx, vz float64) ([2]float64, error) {
	x1 := h1.x
	z1 := h1.z
	x2 := h1.x + h1.dx - vx
	z2 := h1.z + h1.dz - vz
	x3 := h2.x
	z3 := h2.z
	x4 := h2.x + h2.dx - vx
	z4 := h2.z + h2.dz - vz
	if denominator := (z4-z3)*(x2-x1) - (x4-x3)*(z2-z1); denominator == 0 {
		return [2]float64{}, errors.New("invalid denominator")
	} else {
		ua := ((x4-x3)*(z1-z3) - (z4-z3)*(x1-x3)) / denominator
		return [2]float64{x1 + ua*(x2-x1), z1 + ua*(z2-z1)}, nil
	}
}

func findSourceVelocity(in []string, threshold int, rangeX, rangeY, rangeZ int) ([]hailstone, error) {
	var h []hailstone
	if hailstones, err := parseEquations(in); err != nil {
		return h, err
	} else {
		for vz := -rangeZ; vz <= rangeZ; vz++ {
			for vx := -rangeX; vx <= rangeX; vx++ {
				for vy := -rangeY; vy <= rangeY; vy++ {
					skip := false
					count := 0
					for i := 0; i < len(hailstones); i++ {
						for j := i + 1; j < len(hailstones); j++ {
							intXY, eXY := intersectXY(hailstones[i], hailstones[j], float64(vx), float64(vy))
							intYZ, eYZ := intersectYZ(hailstones[i], hailstones[j], float64(vy), float64(vz))
							intXZ, eXZ := intersectXZ(hailstones[i], hailstones[j], float64(vx), float64(vz))
							// Parallel count {
							parallelCount := 0
							if eXY != nil {
								parallelCount++
							}
							if eYZ != nil {
								parallelCount++
							}
							if eXZ != nil {
								parallelCount++
							}
							if parallelCount > 2 {
								skip = true
								break
							}
							matches := 0
							var tI, tJ float64
							if eXY == nil {
								var timeI float64
								if hailstones[i].dx-float64(vx) != 0 {
									timeI = (intXY[0] - hailstones[i].x) / (hailstones[i].dx - float64(vx))
								} else {
									timeI = (intXY[1] - hailstones[i].y) / (hailstones[i].dy - float64(vy))
								}
								var timeJ float64
								if hailstones[j].dx-float64(vx) != 0 {
									timeJ = (intXY[0] - hailstones[j].x) / (hailstones[j].dx - float64(vx))
								} else {
									timeJ = (intXY[1] - hailstones[j].y) / (hailstones[j].dy - float64(vy))
								}
								if tI == 0 {
									tI = timeI
								}
								if tJ == 0 {
									tJ = timeJ
								}
								if timeI > 0 && timeJ > 0 && tI == timeI && tJ == timeJ {
									matches++
								}
							}
							if eYZ == nil {
								var timeI float64
								if hailstones[i].dy-float64(vy) != 0 {
									timeI = (intYZ[0] - hailstones[i].y) / (hailstones[i].dy - float64(vy))
								} else {
									timeI = (intYZ[1] - hailstones[i].z) / (hailstones[i].dz - float64(vz))
								}
								var timeJ float64
								if hailstones[j].dy-float64(vy) != 0 {
									timeJ = (intYZ[0] - hailstones[j].y) / (hailstones[j].dy - float64(vy))
								} else {
									timeJ = (intYZ[1] - hailstones[j].z) / (hailstones[j].dz - float64(vz))
								}
								if tI == 0 {
									tI = timeI
								}
								if tJ == 0 {
									tJ = timeJ
								}
								if timeI > 0 && timeJ > 0 && tI == timeI && tJ == timeJ {
									matches++
								}
							}

							if eXZ == nil {
								var timeI float64
								if hailstones[i].dx-float64(vx) != 0 {
									timeI = (intXZ[0] - hailstones[i].x) / (hailstones[i].dx - float64(vx))
								} else {
									timeI = (intXZ[1] - hailstones[i].z) / (hailstones[i].dz - float64(vz))
								}
								var timeJ float64
								if hailstones[j].dx-float64(vx) != 0 {
									timeJ = (intXZ[0] - hailstones[j].x) / (hailstones[j].dx - float64(vx))
								} else {
									timeJ = (intXZ[1] - hailstones[j].z) / (hailstones[j].dz - float64(vz))
								}
								if tI == 0 {
									tI = timeI
								}
								if tJ == 0 {
									tJ = timeJ
								}
								if timeI > 0 && timeJ > 0 && tI == timeI && tJ == timeJ {
									matches++
								}
							}
							if matches >= 2 {
								count++
							} else {
								skip = true
								break
							}
							if count >= threshold && parallelCount == 0 {
								exist := false
								for _, v := range h {
									if v.x == intXY[0] && v.y == intXY[1] && v.z == intYZ[1] {
										exist = true
										break
									}
								}
								if !exist {
									h = append(h, hailstone{
										x: intXY[0],
										y: intXY[1],
										z: intYZ[1],
									})
								}
							}
						}
						if skip {
							break
						}
					}
				}
			}
		}
	}
	return h, nil
}

func main() {
	abs, _ := filepath.Abs("src/day24/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getCrossedHailstones(output, 200000000000000.0, 400000000000000.0))
	h, _ := findSourceVelocity(output, 20, 500, 500, 500)
	fmt.Println(h[0].x + h[0].y + h[0].z) // 856642398547748
}
