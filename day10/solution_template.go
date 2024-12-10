package day10

import (
	"strconv"
	"strings"
)

// always increases by a height of exactly 1

type pos struct {
	i, j int
}

func PartOne(input string) string {
	trailHeads := make([]pos, 0)
	lines := strings.Split(input, "\n")
	m := make([][]rune, len(lines))
	for i, l := range lines {
		line := make([]rune, len(l))
		for j, b := range l {
			line[j] = b
			if b == '0' {
				trailHeads = append(trailHeads, pos{
					i: i,
					j: j,
				})
			}
		}
		m[i] = line
	}

	total := 0
	for _, h := range trailHeads {
		tops := ToTops(h, &m, make(map[pos]bool))
		total = total + len(tops)

	}
	return strconv.Itoa(total)

}

func ToTops(now pos, mP *[][]rune, visited map[pos]bool) []pos {
	if _, ok := visited[now]; ok {
		return []pos{}
	}
	visited[now] = true
	m := *mP
	if m[now.i][now.j] == '9' {
		return []pos{now}
	}
	neighbors := neighborsF(now, mP, func(a, b int) bool {
		return m[a][b] == m[now.i][now.j]+1
	})
	res := make([]pos, 0)
	for _, neighbor := range neighbors {
		res = append(res, ToTops(neighbor, mP, visited)...)
	}
	return res
}

func legal(m, n, i, j int) bool {
	if i < 0 || i >= m {
		return false
	}
	if j < 0 || j >= n {
		return false
	}
	return true
}
func neighborsF(now pos, mP *[][]rune, f func(i, j int) bool) []pos {
	i := now.i
	j := now.j
	m := *mP
	mapHeight := len(m)
	mapLength := len(m[0])
	res := make([]pos, 0)
	if legal(mapHeight, mapLength, i+1, j) && f(i+1, j) {
		res = append(res, pos{
			i: i + 1,
			j: j,
		})
	}
	if legal(mapHeight, mapLength, i-1, j) && f(i-1, j) {
		res = append(res, pos{
			i: i - 1,
			j: j,
		})
	}
	if legal(mapHeight, mapLength, i, j+1) && f(i, j+1) {
		res = append(res, pos{
			i: i,
			j: j + 1,
		})
	}
	if legal(mapHeight, mapLength, i, j-1) && f(i, j-1) {
		res = append(res, pos{
			i: i,
			j: j - 1,
		})
	}
	return res
}

func PartTwo(input string) string {
	trailHeads := make([]pos, 0)
	lines := strings.Split(input, "\n")
	m := make([][]rune, len(lines))
	for i, l := range lines {
		line := make([]rune, len(l))
		for j, b := range l {
			line[j] = b
			if b == '0' {
				trailHeads = append(trailHeads, pos{
					i: i,
					j: j,
				})
			}
		}
		m[i] = line
	}

	total := 0
	for _, h := range trailHeads {
		tops := ToTops2(h, &m)
		total = total + len(tops)

	}
	return strconv.Itoa(total)
}
func ToTops2(now pos, mP *[][]rune) []pos {
	m := *mP
	if m[now.i][now.j] == '9' {
		return []pos{now}
	}
	neighbors := neighborsF(now, mP, func(a, b int) bool {
		return m[a][b] == m[now.i][now.j]+1
	})
	res := make([]pos, 0)
	for _, neighbor := range neighbors {
		res = append(res, ToTops2(neighbor, mP)...)
	}
	return res
}
