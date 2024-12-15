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

type object2 struct {
	t    objectType
	l, r runeMap.Pos
}

func PartTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	m := make(map[runeMap.Pos]*object2)
	robotPos := runeMap.Pos{}
	lines := strings.Split(parts[0], "\n")
	iMax := len(lines)
	jMax := len(lines[0])
	for i, l := range lines {
		for j, b := range l {
			if b == '@' {
				pos1 := runeMap.Pos{
					I: i,
					J: j * 2,
				}
				robotPos = pos1
				m[pos1] = &object2{t: ROBOT}
			} else if b == '#' {
				pos1 := runeMap.Pos{
					I: i,
					J: j * 2,
				}
				pos2 := runeMap.Pos{
					I: i,
					J: j*2 + 1,
				}
				m[pos1] = &object2{t: ROCK}
				m[pos2] = &object2{t: ROCK}
			} else if b == 'O' {
				pos1 := runeMap.Pos{
					I: i,
					J: j * 2,
				}
				pos2 := runeMap.Pos{
					I: i,
					J: j*2 + 1,
				}
				o := object2{
					t: BOX,
					l: pos1,
					r: pos2,
				}
				m[pos1] = &o
				m[pos2] = &o
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
			//success := tryToMove(&m, dest, ROBOT, move)
			if movable(&m, dest, d) {
				moveIt(&m, dest, &object2{t: ROBOT}, d)
				robotPos = dest
			}
		}
	}
	printMap2(jMax, iMax, &m)
	return strconv.Itoa(count2(&m))
}
func movable(mP *map[runeMap.Pos]*object2, dest runeMap.Pos, d rune) bool {
	var move func(pos runeMap.Pos) runeMap.Pos

	if d == '<' {
		move = left
	} else if d == '^' {
		move = up
	} else if d == 'v' {
		move = down
	} else if d == '>' {
		move = right
	}
	m := *mP
	if v, ok := m[dest]; !ok {
		return true
	} else {
		if v.t == ROCK {
			return false
		}
		if v.t != BOX {
			panic(v.t)
		}
		l, r := v.l, v.r
		pushable := false
		if '>' == d {
			pushable = movable(mP, move(r), d)
		} else if '<' == d {
			pushable = movable(mP, move(l), d)
		} else {
			pushable = movable(mP, move(l), d) && movable(mP, move(r), d)
		}
		if pushable {
			return true
		}
		return false
	}
}

func moveIt(mP *map[runeMap.Pos]*object2, dest runeMap.Pos, o *object2, d rune) {
	var move func(pos runeMap.Pos) runeMap.Pos

	if d == '<' {
		move = left
	} else if d == '^' {
		move = up
	} else if d == 'v' {
		move = down
	} else if d == '>' {
		move = right
	}
	m := *mP
	if v, ok := m[dest]; !ok {
		m[dest] = o
	} else {
		l, r := v.l, v.r
		if '>' == d {
			moveIt(mP, move(dest), v, d)
			m[dest] = o
			if o.r == dest {
				o.r = move(o.r)
				o.l = move(o.l)
			}
		} else if '<' == d {
			moveIt(mP, move(dest), v, d)
			m[dest] = o
			if o.l == dest {
				o.r = move(o.r)
				o.l = move(o.l)
			}
		} else {
			delete(m, v.l)
			delete(m, v.r)
			v.l = move(v.l)
			v.r = move(v.r)
			moveIt(mP, move(l), v, d)
			moveIt(mP, move(r), v, d)
			m[dest] = o
		}
	}
}
func printMap2(ySize, xSize int, m *map[runeMap.Pos]*object2) {
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize*2; j++ {
			if v, ok := (*m)[runeMap.Pos{I: i, J: j}]; ok {
				if v.t == ROCK {
					fmt.Printf("#")
				} else if v.t == BOX {
					if v.l.I == i && v.l.J == j {
						fmt.Printf("[")
					} else if v.r.I == i && v.r.J == j {
						fmt.Printf("]")
					}
				} else if v.t == ROBOT {
					fmt.Printf("@")
				} else {
				}
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()

	}
	fmt.Println()
	fmt.Println()
}
func count2(m *map[runeMap.Pos]*object2) int {
	res := 0
	for k, _ := range *m {
		if v, ok := (*m)[k]; ok {
			if v.t == BOX && v.l.J == k.J {
				res = res + 100*k.I + k.J
			}
		}
	}
	return res
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
