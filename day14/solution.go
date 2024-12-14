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

// PartOne 101 tiles wide and 103 tiles tall,148051800 wrong
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

	xSize, ySize := 101, 103
	allPos := make([]runeMap.Pos, 0, len(robots))
	for _, r := range robots {
		destX, destY := 0, 0
		rowStep := (100 * r.vx % xSize) + r.x
		if rowStep >= xSize {
			destX = rowStep - xSize
		} else if rowStep < 0 {
			rowStep = rowStep + xSize
		} else {
			destX = rowStep
		}

		colStep := (100 * r.vy % ySize) + r.y
		if colStep >= ySize {
			destY = colStep - ySize
		} else if colStep < 0 {
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
	//for j := 0; j < ySize; j++ {
	//	for i := 0; i < xSize; i++ {
	//		v := m[runeMap.Pos{
	//			I: i,
	//			J: j,
	//		}]
	//		if v == 0 {
	//			fmt.Printf(" . ")
	//		} else {
	//
	//			fmt.Printf(" %v ", v)
	//		}
	//	}
	//	fmt.Println()
	//}
	total := one * two * three * four

	return strconv.Itoa(total)
}

func PartTwo(input string) string {
	return ""

}
