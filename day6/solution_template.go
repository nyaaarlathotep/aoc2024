package day6

import (
	"reflect"
	"strconv"
	"strings"
)

type pos struct {
	i, j int
}

func PartOne(input string) string {
	startI, startJ := 0, 0
	lines := strings.Split(input, "\n")
	graph := make([][]rune, len(lines))
	for i, l := range lines {
		line := make([]rune, len(l))
		for j, b := range l {
			line[j] = b
			if b == '^' {
				startI = i
				startJ = j
			}
		}
		graph[i] = line
	}

	dir := up
	m := make(map[pos]bool)
	for {
		m[pos{
			i: startI,
			j: startJ,
		}] = true
		nextI, nextJ := dir(startI, startJ)
		if !legal(graph, nextI, nextJ) {
			break
		}
		if graph[nextI][nextJ] == '#' {
			dir = next(dir)
		} else {
			startI = nextI
			startJ = nextJ
		}
	}

	return strconv.Itoa(len(m))
}

func up(i, j int) (int, int) {
	return i - 1, j
}
func down(i, j int) (int, int) {
	return i + 1, j
}
func left(i, j int) (int, int) {
	return i, j - 1
}
func right(i, j int) (int, int) {
	return i, j + 1
}

func next(f func(int, int) (int, int)) func(int, int) (int, int) {
	fV := reflect.ValueOf(f)
	upV := reflect.ValueOf(up)
	if upV.Pointer() == fV.Pointer() {
		return right
	}
	rightV := reflect.ValueOf(right)
	if rightV.Pointer() == fV.Pointer() {
		return down
	}
	pointer := getFuncPtr(down)
	if pointer == fV.Pointer() {
		return left
	}
	leftV := reflect.ValueOf(left)
	if leftV.Pointer() == fV.Pointer() {
		return up
	}
	panic('!')
}

func getFuncPtr(f func(int, int) (int, int)) uintptr {
	downV := reflect.ValueOf(f)
	pointer := downV.Pointer()
	return pointer
}

func legal(g [][]rune, i, j int) bool {
	if i < 0 || i >= len(g) {
		return false
	}
	if j < 0 || j >= len(g[0]) {
		return false
	}
	return true
}

type posDir struct {
	p   pos
	dir uintptr
}

func PartTwo(input string) string {
	initI, initJ := 0, 0
	lines := strings.Split(input, "\n")
	graph := make([][]rune, len(lines))
	for i, l := range lines {
		line := make([]rune, len(l))
		for j, b := range l {
			line[j] = b
			if b == '^' {
				initI = i
				initJ = j
			}
		}
		graph[i] = line
	}

	startI, startJ := initI, initJ
	dir := up
	m := make(map[pos]bool)
	for {
		m[pos{
			i: startI,
			j: startJ,
		}] = true
		nextI, nextJ := dir(startI, startJ)
		if !legal(graph, nextI, nextJ) {
			break
		}
		if graph[nextI][nextJ] == '#' {
			dir = next(dir)
		} else {
			startI = nextI
			startJ = nextJ
		}
	}
	delete(m, pos{
		i: initI,
		j: initJ,
	})
	res := make([]pos, 0)
	for k, _ := range m {
		graph[k.i][k.j] = '#'

		newM := make(map[posDir]bool)
		startI, startJ = initI, initJ
		dir = up
		for {
			posDirNow := posDir{
				p:   pos{i: startI, j: startJ},
				dir: getFuncPtr(dir),
			}
			if _, ok := newM[posDirNow]; ok {
				res = append(res, pos{
					i: k.i,
					j: k.j,
				})
				break
			} else {
				newM[posDirNow] = true
			}
			nextI, nextJ := dir(startI, startJ)
			if !legal(graph, nextI, nextJ) {
				break
			}
			if graph[nextI][nextJ] == '#' {
				dir = next(dir)
			} else {
				startI = nextI
				startJ = nextJ
			}
		}

		graph[k.i][k.j] = '.'
	}

	return strconv.Itoa(len(res))
}
