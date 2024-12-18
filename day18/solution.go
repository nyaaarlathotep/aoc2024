package day18

import (
	"aoc2024/runeMap"
	"fmt"
	"strconv"
	"strings"
)

const maxI, maxJ = 71, 71
const bytesNum = 1024

func PartOne(input string) string {
	corrupted := make(map[runeMap.Pos]bool)
	for i, cord := range strings.Split(input, "\n") {
		if i >= bytesNum {
			break
		}
		parts := strings.Split(cord, ",")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		corrupted[runeMap.Pos{I: l, J: r}] = true
	}

	now := runeMap.Pos{
		I: 0,
		J: 0,
	}
	end := runeMap.Pos{
		I: maxI - 1,
		J: maxJ - 1,
	}

	minStep := make(map[runeMap.Pos]int)
	outerRange := []runeMap.Pos{now}
	steps := 1
	for len(outerRange) != 0 {
		newOuterRange := make([]runeMap.Pos, 0)
		for _, p := range outerRange {
			neighbors := runeMap.NeighborsWithMNF(p, maxI, maxJ, func(neighborI, neighborJ int) bool {
				pos := runeMap.Pos{I: neighborI, J: neighborJ}
				_, exist := minStep[pos]
				_, corrupt := corrupted[pos]
				return (!exist) && (!corrupt)
			})
			for _, n := range neighbors {
				newOuterRange = append(newOuterRange, n)
				minStep[n] = steps
				if n == end {
					//printMap(maxI, maxJ, minStep, corrupted)
					return strconv.Itoa(steps)
				}
			}
		}

		//printMap(maxI, maxJ, minStep, corrupted)
		outerRange = newOuterRange
		steps++
	}

	return ""

}

func printMap(maxI int, maxJ int, minStep map[runeMap.Pos]int, corrupted map[runeMap.Pos]bool) {
	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			pos := runeMap.Pos{
				I: j,
				J: i,
			}
			if _, ok := minStep[pos]; ok {
				fmt.Printf("O")
			} else if _, ok := corrupted[pos]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func PartTwo(input string) string {
	return ""

}
