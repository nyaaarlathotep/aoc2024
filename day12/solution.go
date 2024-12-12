package day12

import (
	"aoc2024/runeMap"
	"strconv"
	"strings"
)

func PartOne(input string) string {

	lines := strings.Split(input, "\n")
	visited := make([][]bool, len(lines))
	m := make([][]rune, len(lines))
	for i, l := range lines {
		visited[i] = make([]bool, len(l))
		runesLine := make([]rune, len(l))
		for j, r := range l {
			runesLine[j] = r
		}
		m[i] = runesLine
	}

	areas := make([][]runeMap.Pos, 0)
	for i := range m {
		for j := range m[i] {
			nowI, nowJ := i, j
			for !visited[nowI][nowJ] {
				thisArea := make([]runeMap.Pos, 0)
				visit(&visited, &m, runeMap.Pos{
					I: i,
					J: j,
				}, &thisArea)
				areas = append(areas, thisArea)
			}
		}
	}
	total := 0
	for _, a := range areas {
		perimeter := 0
		for _, plot := range a {
			neighbors := runeMap.IllegalNeighborsF(plot, &m, func(neighborI, neighborJ int) bool {
				return m[neighborI][neighborJ] != m[plot.I][plot.J]
			})
			perimeter += len(neighbors)
		}
		now := len(a) * perimeter
		total += now
	}
	return strconv.Itoa(total)
}

func visit(visited *[][]bool, mP *[][]rune, now runeMap.Pos, area *[]runeMap.Pos) {
	*area = append(*area, now)
	(*visited)[now.I][now.J] = true
	m := *mP
	neighbors := runeMap.NeighborsF(now, mP, func(neighborI, neighborJ int) bool {
		return m[neighborI][neighborJ] == m[now.I][now.J] && !(*visited)[neighborI][neighborJ]
	})
	for _, n := range neighbors {

		(*visited)[n.I][n.J] = true
	}
	for _, n := range neighbors {
		visit(visited, mP, n, area)
	}
}

func PartTwo(input string) string {
	return ""

}
