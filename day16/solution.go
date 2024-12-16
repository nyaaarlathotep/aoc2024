package day16

import (
	"aoc2024/runeMap"
	"math"
	"strconv"
	"strings"
)

type posAndDir struct {
	p     runeMap.Pos
	dir   dirInt
	score int
}

type dirInt int

const RIGHT = 1
const DOWN = 2
const LEFT = 3
const UP = 4

var m, n int

func PartOne(input string) string {
	rockMap := make(map[runeMap.Pos]bool)
	s, e := runeMap.Pos{}, runeMap.Pos{}
	lines := strings.Split(input, "\n")
	m = len(lines)
	n = len(lines[0])
	for i, l := range lines {
		for j, b := range l {
			pos := runeMap.Pos{
				I: i,
				J: j,
			}
			if b == 'S' {
				s = pos
				continue
			} else if b == 'E' {
				e = pos
				continue
			} else if b == '#' {
				rockMap[pos] = true
			}
		}
	}
	dirs := make(map[posAndDir]int)
	findWay(&rockMap, &dirs, posAndDir{p: s, dir: RIGHT, score: 0}, e)
	minScore := math.MaxInt
	minScoreR := dirs[posAndDir{p: e, dir: RIGHT, score: 0}]
	if minScoreR != 0 {
		minScore = min(minScoreR, minScoreR)
	}
	minScoreD := dirs[posAndDir{p: e, dir: DOWN, score: 0}]
	if minScoreD != 0 {
		minScore = min(minScoreD, minScoreR)
	}
	minScoreL := dirs[posAndDir{p: e, dir: LEFT, score: 0}]
	if minScoreL != 0 {
		minScore = min(minScoreL, minScoreR)
	}
	minScoreU := dirs[posAndDir{p: e, dir: UP, score: 0}]
	if minScoreU != 0 {
		minScore = min(minScoreU, minScoreR)
	}
	return strconv.Itoa(minScore)
}

func findWay(rockMap *map[runeMap.Pos]bool, minScoreMap *map[posAndDir]int, start posAndDir, end runeMap.Pos) {
	nowPoses := []posAndDir{start}
	for len(nowPoses) != 0 {
		p := nowPoses[0]
		nowPoses = nowPoses[1:]
		if p.p == end {
			continue
		}
		neighbors := transF(p, func(neighborI, neighborJ int) bool {
			return !(*rockMap)[runeMap.Pos{
				I: neighborI,
				J: neighborJ,
			}]
		})
		for _, neighbor := range neighbors {
			k := neighbor
			k.score = 0
			if v, ok := (*minScoreMap)[k]; ok {
				if neighbor.score < v {
					(*minScoreMap)[k] = neighbor.score
					nowPoses = append(nowPoses, neighbor)
				}
			} else {
				(*minScoreMap)[k] = neighbor.score
				nowPoses = append(nowPoses, neighbor)
			}
		}
	}
}
func transF(now posAndDir, f func(neighborI, neighborJ int) bool) []posAndDir {
	i := now.p.I
	j := now.p.J
	mapHeight := m
	mapLength := n
	res := make([]posAndDir, 0)
	if now.dir == UP {
		l, r := now, now
		l.dir = LEFT
		l.score = l.score + 1000
		r.dir = RIGHT
		r.score = r.score + 1000
		res = append(res, l)
		res = append(res, r)
		if legal(mapHeight, mapLength, i-1, j) && f(i-1, j) {
			res = append(res, posAndDir{
				p:     runeMap.Pos{I: i - 1, J: j},
				dir:   UP,
				score: now.score + 1,
			})
		}
	} else if now.dir == RIGHT {
		l, r := now, now
		l.dir = UP
		l.score = l.score + 1000
		r.dir = DOWN
		r.score = r.score + 1000
		res = append(res, l)
		res = append(res, r)
		if legal(mapHeight, mapLength, i, j+1) && f(i, j+1) {
			res = append(res, posAndDir{
				p:     runeMap.Pos{I: i, J: j + 1},
				dir:   RIGHT,
				score: now.score + 1,
			})
		}
	} else if now.dir == DOWN {
		l, r := now, now
		l.dir = LEFT
		l.score = l.score + 1000
		r.dir = RIGHT
		r.score = r.score + 1000
		res = append(res, l)
		res = append(res, r)
		if legal(mapHeight, mapLength, i+1, j) && f(i+1, j) {
			res = append(res, posAndDir{
				p:     runeMap.Pos{I: i + 1, J: j},
				dir:   DOWN,
				score: now.score + 1,
			})
		}
	} else if now.dir == LEFT {
		l, r := now, now
		l.dir = UP
		l.score = l.score + 1000
		r.dir = DOWN
		r.score = r.score + 1000
		res = append(res, l)
		res = append(res, r)
		if legal(mapHeight, mapLength, i, j-1) && f(i, j-1) {
			res = append(res, posAndDir{
				p:     runeMap.Pos{I: i, J: j - 1},
				dir:   LEFT,
				score: now.score + 1,
			})
		}
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
func PartTwo(input string) string {
	return ""

}
