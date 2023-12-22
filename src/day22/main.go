package main

import (
	"advent2023/pkg/file"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type bricks struct {
	sx, sy, sz int
	dx, dy, dz int
	idx        int
}

func parseInput(in []string) (brickList []bricks, maxX, maxY, maxZ int, err error) {
	var (
		sx, sy, sz int
		dx, dy, dz int
	)
	for i, l := range in {
		sd := strings.Split(l, "~")
		if len(sd) != 2 {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		sxyz := strings.Split(sd[0], ",")
		if len(sxyz) != 3 {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		dxyz := strings.Split(sd[1], ",")
		if len(dxyz) != 3 {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if sx, err = strconv.Atoi(sxyz[0]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if sy, err = strconv.Atoi(sxyz[1]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if sz, err = strconv.Atoi(sxyz[2]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if dx, err = strconv.Atoi(dxyz[0]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if dy, err = strconv.Atoi(dxyz[1]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		if dz, err = strconv.Atoi(dxyz[2]); err != nil {
			return brickList, maxX, maxY, maxZ, errors.New("invalid input")
		}
		brickList = append(brickList, bricks{
			sx:  sx,
			sy:  sy,
			sz:  sz,
			dx:  dx,
			dy:  dy,
			dz:  dz,
			idx: i + 1,
		})
		if sx > maxX {
			maxX = sx
		}
		if dx > maxX {
			maxX = dx
		}
		if sy > maxY {
			maxY = sy
		}
		if dy > maxY {
			maxY = dy
		}
		if sz > maxZ {
			maxZ = sz
		}
		if dz > maxZ {
			maxZ = dz
		}
	}
	return brickList, maxX, maxY, maxZ, err
}

func generateMaps(brickList []bricks, maxX, maxY, maxZ int) (mxz, myz [][]int) {
	// Create empty maps
	mxz = make([][]int, maxZ+1)
	myz = make([][]int, maxZ+1)
	for z := 0; z < maxZ+1; z++ {
		mxz[z] = make([]int, maxX+1)
		myz[z] = make([]int, maxY+1)
	}
	// Set bricks on maps
	for _, brick := range brickList {
		for z := brick.sz; z <= brick.dz; z++ {
			for x := brick.sx; x <= brick.dx; x++ {
				if mxz[z][x] == 0 {
					mxz[z][x] = brick.idx
				} else {
					mxz[z][x] = 99999
				}
			}
			for y := brick.sy; y <= brick.dy; y++ {
				if myz[z][y] == 0 {
					myz[z][y] = brick.idx
				} else {
					myz[z][y] = 99999
				}
			}
		}
	}
	return mxz, myz
}

func plotMaps(mxz, myz [][]int) {
	for z := len(mxz) - 1; z >= 0; z-- {
		fmt.Printf("zx[%d] -> ", z)
		for x := 0; x < len(mxz[z]); x++ {
			if mxz[z][x] == 99999 {
				fmt.Printf("?, ")
			} else {
				fmt.Printf("%d, ", mxz[z][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	for z := len(mxz) - 1; z >= 0; z-- {
		fmt.Printf("zy[%d] -> ", z)
		for y := 0; y < len(myz[z]); y++ {
			if myz[z][y] == 99999 {
				fmt.Printf("?, ")
			} else {
				fmt.Printf("%d, ", myz[z][y])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getSettledBricks(brickList []bricks) []bricks {
	for {
		finished := true
		sort.Slice(brickList, func(i, j int) bool {
			return brickList[i].sz < brickList[j].sz
		})
		for idx, brick := range brickList {
			if brick.sz > 0 {
				var touch bool
				for j := 0; j < idx; j++ {
					touch = false
					if brickList[j].dz == brick.sz-1 {
						if brick.sx <= brickList[j].dx && brick.dx >= brickList[j].sx && brick.sy <= brickList[j].dy && brick.dy >= brickList[j].sy {
							touch = true
							break
						}
					}
				}
				if !touch {
					finished = false
					brickList[idx].sz--
					brickList[idx].dz--
				}
			}
		}
		if finished {
			break
		}
	}
	return brickList
}

func countSafelyDesintegratedBrisks(brickList []bricks) int {
	count := 0
	for idx := 0; idx < len(brickList); idx++ {
		copiedSlice := make([]bricks, len(brickList))
		copy(copiedSlice, brickList)
		copiedSlice = append(copiedSlice[:idx], copiedSlice[idx+1:]...)
		newBrickList := make([]bricks, len(copiedSlice))
		copy(newBrickList, copiedSlice)
		newBrickList = getSettledBricks(newBrickList)
		diff := false
		for i := 0; i < len(copiedSlice); i++ {
			for j := 0; j < len(newBrickList); j++ {
				if copiedSlice[i].idx == newBrickList[j].idx {
					if copiedSlice[i].sx != newBrickList[j].sx ||
						copiedSlice[i].sy != newBrickList[j].sy ||
						copiedSlice[i].sz != newBrickList[j].sz ||
						copiedSlice[i].dx != newBrickList[j].dx ||
						copiedSlice[i].dy != newBrickList[j].dy ||
						copiedSlice[i].dz != newBrickList[j].dz {
						diff = true
					}
					break
				}
			}
		}
		if !diff {
			count++
		}
	}
	return count
}

func getSafelyDesintegratedBrisks(in []string) (int, error) {
	if brickList, maxX, maxY, maxZ, err := parseInput(in); err != nil {
		return 0, err
	} else {
		fmt.Println(maxX, maxY, maxZ)
		//mxz, myz := generateMaps(brickList, maxX, maxY, maxZ)
		//plotMaps(mxz, myz)
		brickList = getSettledBricks(brickList)
		//mxz, myz = generateMaps(brickList, maxX, maxY, maxZ)
		//fmt.Println("\r\n finished")
		//plotMaps(mxz, myz)
		return countSafelyDesintegratedBrisks(brickList), nil
	}
}

func getFallen(brickList []bricks) int {
	diff := 0
	for idx := 0; idx < len(brickList); idx++ {
		copiedSlice := make([]bricks, len(brickList))
		copy(copiedSlice, brickList)
		copiedSlice = append(copiedSlice[:idx], copiedSlice[idx+1:]...)
		newBrickList := make([]bricks, len(copiedSlice))
		copy(newBrickList, copiedSlice)
		newBrickList = getSettledBricks(newBrickList)
		for i := 0; i < len(copiedSlice); i++ {
			for j := 0; j < len(newBrickList); j++ {
				if copiedSlice[i].idx == newBrickList[j].idx {
					if copiedSlice[i].sx != newBrickList[j].sx ||
						copiedSlice[i].sy != newBrickList[j].sy ||
						copiedSlice[i].sz != newBrickList[j].sz ||
						copiedSlice[i].dx != newBrickList[j].dx ||
						copiedSlice[i].dy != newBrickList[j].dy ||
						copiedSlice[i].dz != newBrickList[j].dz {
						diff++
					}
					break
				}
			}
		}
	}
	return diff
}
func getTotalFallen(in []string) (int, error) {
	if brickList, maxX, maxY, maxZ, err := parseInput(in); err != nil {
		return 0, err
	} else {
		fmt.Println(maxX, maxY, maxZ)
		//mxz, myz := generateMaps(brickList, maxX, maxY, maxZ)
		//plotMaps(mxz, myz)
		brickList = getSettledBricks(brickList)
		//mxz, myz = generateMaps(brickList, maxX, maxY, maxZ)
		//fmt.Println("\r\n finished")
		//plotMaps(mxz, myz)
		return getFallen(brickList), nil
	}
}

func main() {
	abs, _ := filepath.Abs("src/day22/input.txt")
	output, _ := file.ReadInput(abs)
	fmt.Println(getSafelyDesintegratedBrisks(output))
	fmt.Println(getTotalFallen(output))
}
