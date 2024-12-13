package day13

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type remainXY struct {
	x, y int
}

type tokenState struct {
	a, b int
}

// PartOne each button would need to be pressed no more than 100 times
func PartOne(input string) string {
	parts := strings.Split(input, "\n\n")
	re1 := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	re2 := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	aTokens := 3
	bTokens := 1
	totalMin := 0
	for _, p := range parts {
		var ax, ay, bx, by, rx, ry int
		lines := strings.Split(p, "\n")
		matches := re1.FindAllStringSubmatch(lines[0], -1)
		for _, match := range matches {
			ax, _ = strconv.Atoi(match[1])
			ay, _ = strconv.Atoi(match[2])
		}
		matches = re1.FindAllStringSubmatch(lines[1], -1)
		for _, match := range matches {
			bx, _ = strconv.Atoi(match[1])
			by, _ = strconv.Atoi(match[2])
		}
		matches = re2.FindAllStringSubmatch(lines[2], -1)
		for _, match := range matches {
			rx, _ = strconv.Atoi(match[1])
			ry, _ = strconv.Atoi(match[2])
		}

		var getCount func(int, int)
		succeedTokens := make([]tokenState, 0)
		getCount = func(remainX int, remainY int) {
			var dp [][]remainXY
			dp = append(dp, []remainXY{{
				x: rx,
				y: ry,
			}})
			for i := 1; i <= 100; i++ {
				remain := dp[i-1][0]
				xy := remainXY{
					x: remain.x - ax,
					y: remain.y - ay,
				}
				if xy.x == 0 && xy.y == 0 {
					succeedTokens = append(succeedTokens, tokenState{
						a: i,
						b: 0,
					})
					break
				}
				if xy.x <= 0 || xy.y <= 0 {
					break
				}
				dp = append(dp, []remainXY{xy})
			}

			for j := 0; j < len(dp); j++ {
				for i := 1; i <= 100; i++ {
					remain := dp[j][i-1]
					xy := remainXY{
						x: remain.x - bx,
						y: remain.y - by,
					}
					if xy.x == 0 && xy.y == 0 {
						succeedTokens = append(succeedTokens, tokenState{
							a: j,
							b: i,
						})
						break
					}
					if xy.x <= 0 || xy.y <= 0 {
						break
					}
					dp[j] = append(dp[j], xy)
				}
			}

		}
		getCount(rx, ry)
		minTokens := math.MaxInt
		for _, t := range succeedTokens {
			totalTokens := t.a*aTokens + t.b*bTokens
			if totalTokens < minTokens {
				minTokens = totalTokens
			}
		}
		if minTokens != math.MaxInt {
			totalMin += minTokens
		}
	}
	return strconv.Itoa(totalMin)
}

func PartTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	re1 := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	re2 := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	totalMin := 0
	for _, p := range parts {
		var ax, ay, bx, by, rx, ry int
		lines := strings.Split(p, "\n")
		matches := re1.FindAllStringSubmatch(lines[0], -1)
		for _, match := range matches {
			ax, _ = strconv.Atoi(match[1])
			ay, _ = strconv.Atoi(match[2])
		}
		matches = re1.FindAllStringSubmatch(lines[1], -1)
		for _, match := range matches {
			bx, _ = strconv.Atoi(match[1])
			by, _ = strconv.Atoi(match[2])
		}
		matches = re2.FindAllStringSubmatch(lines[2], -1)
		for _, match := range matches {
			rx, _ = strconv.Atoi(match[1])
			ry, _ = strconv.Atoi(match[2])
		}
		rx += 10000000000000
		ry += 10000000000000
		nbX := bx * ay
		nbY := by * ax
		npX := rx * ay
		npY := ry * ax

		B := abs(nbX - nbY)
		P := abs(npX - npY)

		if P%B == 0 {
			b := P / B
			if (rx-(b*bx))%ax == 0 {
				a := (rx - (b * bx)) / ax
				totalMin += (a * 3) + b
			}
		}
	}
	return strconv.Itoa(totalMin)
}

var abs = func(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}
