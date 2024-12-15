package day15

import (
	"aoc2024/runeMap"
	"fmt"
	"strconv"
	"strings"
)

type objectType int

const (
	BOX   = 2
	ROCK  = 1
	ROBOT = 3
)

func PartOne(input string) string {
	parts := strings.Split(input, "\n\n")
	m := make(map[runeMap.Pos]objectType)
	robotPos := runeMap.Pos{}
	lines := strings.Split(parts[0], "\n")
	iMax := len(lines)
	jMax := len(lines[0])
	for i, l := range lines {
		for j, b := range l {
			if b == '@' {
				pos := runeMap.Pos{
					I: i,
					J: j,
				}
				robotPos = pos
				m[pos] = ROBOT
			} else if b == '#' {
				m[runeMap.Pos{
					I: i,
					J: j,
				}] = ROCK
			} else if b == 'O' {
				m[runeMap.Pos{
					I: i,
					J: j,
				}] = BOX
			}
		}
	}

	for _, l := range strings.Split(parts[1], "\n") {
		var move func(pos runeMap.Pos) runeMap.Pos
		for _, d := range l {
			if d == '<' {
				move = left
			} else if d == '^' {
				move = up
			} else if d == 'v' {
				move = down
			} else if d == '>' {
				move = right
			} else {
				panic(d)
			}
			delete(m, robotPos)
			dest := move(robotPos)
			success := tryToMove(&m, dest, ROBOT, move)
			if success {
				robotPos = dest
			}
		}
	}
	printMap(jMax, iMax, &m)
	return strconv.Itoa(count(&m))
}

func count(m *map[runeMap.Pos]objectType) int {
	res := 0
	for k, v := range *m {
		if v == BOX {
			res = res + 100*k.I + k.J
		}
	}
	return res
}

func up(pos runeMap.Pos) runeMap.Pos {
	return runeMap.Pos{
		I: pos.I - 1,
		J: pos.J,
	}
}

func down(pos runeMap.Pos) runeMap.Pos {
	return runeMap.Pos{
		I: pos.I + 1,
		J: pos.J,
	}
}

func right(pos runeMap.Pos) runeMap.Pos {
	return runeMap.Pos{
		I: pos.I,
		J: pos.J + 1,
	}
}

func left(pos runeMap.Pos) runeMap.Pos {
	return runeMap.Pos{
		I: pos.I,
		J: pos.J - 1,
	}
}

func tryToMove(mP *map[runeMap.Pos]objectType, dest runeMap.Pos, coming objectType, f func(pos runeMap.Pos) runeMap.Pos) bool {
	m := *mP
	if v, ok := m[dest]; !ok {
		m[dest] = coming
		return true
	} else {
		if v == ROCK {
			return false
		}
		pushNext := tryToMove(mP, f(dest), v, f)
		if pushNext {
			m[dest] = coming
			return true
		}
		return false
	}
}

func PartTwo(input string) string {
	return ""

}

func printMap(ySize, xSize int, m *map[runeMap.Pos]objectType) {
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			v := (*m)[runeMap.Pos{
				I: i,
				J: j,
			}]
			if v == ROCK {
				fmt.Printf("#")
			} else if v == BOX {
				fmt.Printf("O")
			} else if v == ROBOT {
				fmt.Printf("@")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()

	}
	fmt.Println()
	fmt.Println()
}
