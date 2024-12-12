package day12

import (
	"aoc2024/runeMap"
	"sort"
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
		allNeighbors := make([]fence, 0)
		for _, plot := range a {
			neighbors := runeMap.IllegalNeighborsF(plot, &m, func(neighborI, neighborJ int) bool {
				return m[neighborI][neighborJ] != m[plot.I][plot.J]
			})
			for _, n := range neighbors {
				allNeighbors = append(allNeighbors, fence{
					inner: plot,
					outer: n,
				})
			}
		}
		perimeter := getPerimeter(allNeighbors)
		//fmt.Printf("%c: %d * %d\n", m[a[0].I][a[0].J], len(a), perimeter)
		now := len(a) * perimeter
		total += now
	}
	return strconv.Itoa(total)
}

type fence struct {
	inner, outer runeMap.Pos
}

func getPerimeter(neighbors []fence) int {
	added := make(map[fence]bool)
	groupCount := 0
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].outer.I < neighbors[j].outer.I || neighbors[i].outer.J < neighbors[j].outer.J
	})
	for _, f := range neighbors {
		col := false
		if f.inner.I != f.outer.I {
			col = false
		} else {
			col = true
		}
		if col {
			up := fence{
				inner: runeMap.Pos{
					I: f.inner.I - 1,
					J: f.inner.J,
				},
				outer: runeMap.Pos{
					I: f.outer.I - 1,
					J: f.outer.J,
				},
			}
			down := fence{
				inner: runeMap.Pos{
					I: f.inner.I + 1,
					J: f.inner.J,
				},
				outer: runeMap.Pos{
					I: f.outer.I + 1,
					J: f.outer.J,
				},
			}
			if !added[up] && !added[down] {
				groupCount++
			}

		} else {
			left := fence{
				inner: runeMap.Pos{
					I: f.inner.I,
					J: f.inner.J - 1,
				},
				outer: runeMap.Pos{
					I: f.outer.I,
					J: f.outer.J - 1,
				},
			}
			right := fence{
				inner: runeMap.Pos{
					I: f.inner.I,
					J: f.inner.J + 1,
				},
				outer: runeMap.Pos{
					I: f.outer.I,
					J: f.outer.J + 1,
				},
			}
			if !added[left] && !added[right] {
				groupCount++
			}
		}
		added[f] = true
	}
	return groupCount
}
