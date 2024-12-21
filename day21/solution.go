package day21

import (
	"aoc2024/runeMap"
	"math"
	"slices"
	"strconv"
	"strings"
)

var letterMap = map[rune]runeMap.Pos{
	'7': {I: 0, J: 0},
	'8': {I: 0, J: 1},
	'9': {I: 0, J: 2},
	'4': {I: 1, J: 0},
	'5': {I: 1, J: 1},
	'6': {I: 1, J: 2},
	'1': {I: 2, J: 0},
	'2': {I: 2, J: 1},
	'3': {I: 2, J: 2},
	'0': {I: 3, J: 1},
	'a': {I: 3, J: 2},
	'A': {I: 0, J: 2},
	'>': {I: 1, J: 2},
	'^': {I: 0, J: 1},
	'<': {I: 1, J: 0},
	'v': {I: 1, J: 1},
}

var dir = map[byte][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

type state struct {
	s     string
	depth int
	lim   int
}

func PartOne(input string) string {
	total := 0
	m := make(map[state]int)
	for _, l := range strings.Split(input, "\n") {
		numR := make([]rune, 0)
		for _, r := range l {
			if r >= '0' && r <= '9' {
				numR = append(numR, r)
			}
		}
		numLine, _ := strconv.Atoi(string(numR))
		s := state{
			s:     l,
			depth: 0,
			lim:   2,
		}
		length := minLen(s, m)
		total = total + length*numLine
	}

	return strconv.Itoa(total)
}

func minLen(s state, m map[state]int) int {
	if v, ok := m[s]; ok {
		return v
	}
	var gap runeMap.Pos
	if s.depth != 0 {
		gap = runeMap.Pos{}
	} else {
		gap = runeMap.Pos{I: 3, J: 0}
	}
	var now runeMap.Pos
	if s.depth != 0 {
		now = letterMap['A']
	} else {
		now = letterMap['a']
	}

	length := 0
	for _, c := range s.s {
		if s.depth == 0 && c == 'A' {
			c = 'a'
		}

		nextPos := letterMap[c]
		allMoves := getAllPossibleMove(gap, now, nextPos)
		if s.depth == s.lim {
			move := allMoves[0]
			slices.Reverse(move)
			toMove := string(move) + "A"
			length = length + len(toMove)
		} else {
			thisMinLen := math.MaxInt
			for _, move := range allMoves {
				slices.Reverse(move)
				toMove := string(move) + "A"
				s2 := state{
					s:     toMove,
					depth: s.depth + 1,
					lim:   s.lim,
				}
				remainMin := minLen(s2, m)
				thisMinLen = min(thisMinLen, remainMin)
			}
			length = length + thisMinLen
		}
		now = nextPos

	}
	m[s] = length
	return length
}

func getAllPossibleMove(avoid runeMap.Pos, now, end runeMap.Pos) [][]rune {
	if now == avoid {
		return nil
	}
	if now == end {
		return [][]rune{{}}
	}
	yMove, xMove := end.I-now.I, end.J-now.J
	var next runeMap.Pos
	var thisMove rune
	move := true
	res := make([][]rune, 0)
	if xMove < 0 {
		next = runeMap.Pos{I: now.I, J: now.J - 1}
		thisMove = '<'
	} else if xMove > 0 {
		next = runeMap.Pos{I: now.I, J: now.J + 1}
		thisMove = '>'
	} else {
		move = false
	}
	if move {
		if next != avoid {
			allPoss := getAllPossibleMove(avoid, next, end)
			if allPoss != nil {
				for _, route := range allPoss {
					res = append(res, append(route, thisMove))
				}
			}
		}
	}

	move = true
	if yMove < 0 {
		next = runeMap.Pos{I: now.I - 1, J: now.J}
		thisMove = '^'
	} else if yMove > 0 {
		next = runeMap.Pos{I: now.I + 1, J: now.J}
		thisMove = 'v'
	} else {
		move = false
	}
	if move {
		if next != avoid {
			allPoss := getAllPossibleMove(avoid, next, end)
			if allPoss != nil {
				for _, route := range allPoss {
					res = append(res, append(route, thisMove))
				}
			}
		}
	}
	if len(res) == 0 {
		panic("!")
	}
	return res
}

func PartTwo(input string) string {
	total := 0
	m := make(map[state]int)
	for _, l := range strings.Split(input, "\n") {
		numR := make([]rune, 0)
		for _, r := range l {
			if r >= '0' && r <= '9' {
				numR = append(numR, r)
			}
		}
		numLine, _ := strconv.Atoi(string(numR))
		s := state{
			s:     l,
			depth: 0,
			lim:   25,
		}
		length := minLen(s, m)
		total = total + length*numLine
	}

	return strconv.Itoa(total)
}
