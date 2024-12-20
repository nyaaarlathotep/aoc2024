package day20

import (
	"aoc2024/runeMap"
	"fmt"
	"strings"
)

func PartOne(input string) string {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	stones := make(map[runeMap.Pos]bool)
	var s, e runeMap.Pos
	for i, l := range lines {
		for j, b := range l {
			if b == '#' {
				stones[runeMap.Pos{I: i, J: j}] = true
			} else if b == 'S' {
				s = runeMap.Pos{I: i, J: j}
			} else if b == 'E' {
				e = runeMap.Pos{I: i, J: j}
			}
		}
	}

	cheatMap := make(map[int][]runeMap.Pos)
	length := oneBugChance(stones, m, n, []runeMap.Pos{s}, e, cheatMap)

	total := 0
	for k, v := range cheatMap {
		if k == -1 {
			continue
		}
		if length-k >= 100 {
			total += len(v)
		}
		fmt.Printf("%v: %v\n", len(v), length-k)
	}
	fmt.Printf("%v\n", length)
	return ""
}

func oneBugChance(stoneMap map[runeMap.Pos]bool, m, n int,
	nowPoses []runeMap.Pos, e runeMap.Pos, cheatMap map[int][]runeMap.Pos) int {
	stepsMap := make(map[runeMap.Pos]int)
	stepsNow := 1
	for len(nowPoses) != 0 {
		nextPoses := make([]runeMap.Pos, 0)
		for _, p := range nowPoses {
			neighbors := runeMap.NeighborsWithMNF(p, m, n, func(neighborI, neighborJ int) bool {
				pos := runeMap.Pos{I: neighborI, J: neighborJ}
				_, exist := stoneMap[pos]
				_, stepped := stepsMap[pos]
				return (!exist) && (!stepped)
			})

			for _, neighbor := range neighbors {
				nextPoses = append(nextPoses, neighbor)
				stepsMap[neighbor] = stepsNow
				if neighbor == e {
					return stepsNow
				}
			}

			bugNeighbors := runeMap.NeighborsWithMNF(p, m, n, func(neighborI, neighborJ int) bool {
				pos := runeMap.Pos{I: neighborI, J: neighborJ}
				_, exist := stoneMap[pos]
				return exist
			})
			for _, neighbor := range bugNeighbors {
				newStepsMap := make(map[runeMap.Pos]int, len(stepsMap)+1)
				for k, v := range stepsMap {
					newStepsMap[k] = v
				}
				newStepsMap[neighbor] = stepsNow
				cheatMinRes := walk(newStepsMap, stoneMap, m, n, []runeMap.Pos{neighbor}, e, stepsNow+1)
				cheatMap[cheatMinRes] = append(cheatMap[cheatMinRes], neighbor)
			}

		}
		nowPoses = nextPoses
		stepsNow++
	}
	return -1
}
func walk(stepsMap map[runeMap.Pos]int, stoneMap map[runeMap.Pos]bool, m, n int, nowPoses []runeMap.Pos, e runeMap.Pos, stepsNow int) int {
	for len(nowPoses) != 0 {
		nextPoses := make([]runeMap.Pos, 0)
		for _, p := range nowPoses {
			neighbors := runeMap.NeighborsWithMNF(p, m, n, func(neighborI, neighborJ int) bool {
				pos := runeMap.Pos{I: neighborI, J: neighborJ}
				_, exist := stoneMap[pos]
				if exist {
					return false
				}
				_, stepped := stepsMap[pos]
				return (!exist) && (!stepped)
			})
			for _, neighbor := range neighbors {
				nextPoses = append(nextPoses, neighbor)
				stepsMap[neighbor] = stepsNow
				if neighbor == e {
					return stepsNow
				}
			}
		}
		nowPoses = nextPoses
		stepsNow++
	}
	return -1
}
func PartTwo(input string) string {
	return ""

}
