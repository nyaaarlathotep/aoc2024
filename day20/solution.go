package day20

import (
	"aoc2024/runeMap"
	"fmt"
	"strconv"
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
	length := oneBugChance(stones, m, n, []runeMap.Pos{s}, s, e, cheatMap)

	total := 0
	for k, v := range cheatMap {
		if k == -1 {
			continue
		}
		if length-k >= 100 {
			total += len(v)
		}
		//fmt.Printf("%v: %v\n", len(v), length-k)
	}
	//fmt.Printf("%v\n", length)
	//fmt.Println(cheatMap[length-2])
	return strconv.Itoa(total)
}

func oneBugChance(stoneMap map[runeMap.Pos]bool, m, n int,
	nowPoses []runeMap.Pos, s, e runeMap.Pos, cheatMap map[int][]runeMap.Pos) int {
	stepsMap := make(map[runeMap.Pos]int)
	stepsMap[nowPoses[0]] = 0
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
				_, stepped := stepsMap[pos]
				return !stepped
			})
			for _, bugPos := range bugNeighbors {
				newStepsMap := make(map[runeMap.Pos]int, len(stepsMap)+1)
				for k, v := range stepsMap {
					newStepsMap[k] = v
				}
				newStepsMap[bugPos] = stepsNow
				cheatMinRes := walk(newStepsMap, stoneMap, m, n, []runeMap.Pos{bugPos}, s, e, stepsNow+1)
				cheatMap[cheatMinRes] = append(cheatMap[cheatMinRes], bugPos)
			}

		}
		nowPoses = nextPoses
		stepsNow++
	}
	return -1
}
func walk(stepsMap map[runeMap.Pos]int, stoneMap map[runeMap.Pos]bool, m, n int, nowPoses []runeMap.Pos, s, e runeMap.Pos, stepsNow int) int {
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
					//if stepsNow == 82 {
					//	printMap(m, n, stepsMap, stoneMap, s, e)
					//}
					return stepsNow
				}
			}
		}
		nowPoses = nextPoses
		stepsNow++
	}
	return -1
}

func printMap(maxI int, maxJ int, minStep map[runeMap.Pos]int, corrupted map[runeMap.Pos]bool, s, e runeMap.Pos) {
	for j := 0; j < maxJ; j++ {
		for i := 0; i < maxI; i++ {
			pos := runeMap.Pos{
				I: j,
				J: i,
			}
			if pos == s {
				fmt.Printf("S")
			} else if pos == e {
				fmt.Printf("E")
			} else if _, ok := minStep[pos]; ok {
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
	SolvePartTwo("/home/konomi/codes/go/aoc2024/day20/input")
	return ""
}

func twentyBugChance(stoneMap map[runeMap.Pos]bool, m, n int,
	nowPoses []runeMap.Pos, s, e runeMap.Pos, cheatMap map[int][]runeMap.Pos) int {
	stepsMap := make(map[runeMap.Pos]int)
	stepsMap[nowPoses[0]] = 0
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
			goForBug(stoneMap, m, n, p, stepsMap, stepsNow, s, e, cheatMap)
		}
		nowPoses = nextPoses
		stepsNow++
	}
	return -1
}

func goForBug(stoneMap map[runeMap.Pos]bool, m int, n int, p runeMap.Pos, stepsMap map[runeMap.Pos]int, stepsNow int,
	s runeMap.Pos, e runeMap.Pos, cheatMap map[int][]runeMap.Pos) {

	nextNeighbors := make([]runeMap.Pos, 0)
	for i := 0; i < 19; i++ {
		possibleNeighbors := runeMap.NeighborsWithMNF(p, m, n, func(neighborI, neighborJ int) bool {
			pos := runeMap.Pos{I: neighborI, J: neighborJ}
			_, stepped := stepsMap[pos]
			return !stepped
		})

		noNeedForCheat := make([]runeMap.Pos, 0)
		for _, neighbor := range possibleNeighbors {
			nextNeighbors = append(nextNeighbors, neighbor)
			stepsMap[neighbor] = stepsNow
			if neighbor == e {
				noNeedForCheat = append(noNeedForCheat, neighbor)
			}
		}

	}

}
