package day14

import (
	"aoc2024/runeMap"
	"regexp"
	"strconv"
	"strings"
)

type robot struct {
	x, y   int
	vx, vy int
}

// PartOne 101 tiles wide and 103 tiles tall,148051800 wrong, 222208000
func PartOne(input string) string {
	re := regexp.MustCompile(`-?\d+`)
	lines := strings.Split(input, "\n")
	robots := make([]robot, 0, len(lines))
	for _, l := range lines {
		matches := re.FindAllString(l, -1)
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		vx, _ := strconv.Atoi(matches[2])
		vy, _ := strconv.Atoi(matches[3])
		r := robot{
			x:  x,
			y:  y,
			vx: vx,
			vy: vy,
		}
		robots = append(robots, r)
	}

	xSize, ySize, seconds := 101, 103, 100
	allPos := make([]runeMap.Pos, 0, len(robots))
	for _, r := range robots {
		destX, destY := 0, 0
		rowStep := (seconds*r.vx + r.x) % xSize
		if rowStep < 0 {
			destX = rowStep + xSize
		} else {
			destX = rowStep
		}

		colStep := (seconds*r.vy + r.y) % ySize
		if colStep < 0 {
			destY = colStep + ySize
		} else {
			destY = colStep
		}
		allPos = append(allPos, runeMap.Pos{
			I: destX,
			J: destY,
		})
	}

	one, two, three, four := 0, 0, 0, 0
	midX := xSize / 2
	midY := ySize / 2

	m := make(map[runeMap.Pos]int)
	for _, p := range allPos {
		m[p] = m[p] + 1
		if p.I < midX && p.J < midY {
			one++
		} else if p.I > midX && p.J < midY {
			two++
		} else if p.I < midX && p.J > midY {
			three++
		} else if p.I > midX && p.J > midY {
			four++
		}
	}
	total := one * two * three * four

	return strconv.Itoa(total)
}

func PartTwo(input string) string {
	return ""

}
