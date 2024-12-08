package day8

import (
	"strconv"
	"strings"
)

type pos struct {
	name rune
	i, j int
}

func PartOne(input string) string {
	m := make(map[rune][]pos)

	lines := strings.Split(input, "\n")
	mapHeight := len(lines)
	mapLength := len(lines[0])
	for i, l := range lines {
		for j, b := range l {
			if b != '.' {
				if v, ok := m[b]; ok {
					m[b] = append(v, pos{
						name: b,
						i:    i,
						j:    j,
					})
				} else {
					m[b] = []pos{{
						name: b,
						i:    i,
						j:    j,
					}}
				}
			}
		}
	}

	uniquePoses := make(map[pos]bool)
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			one := v[i]
			for j := i + 1; j < len(v); j++ {
				two := v[j]
				i1 := one.i - (two.i - one.i)
				j1 := one.j - (two.j - one.j)
				if legal(mapHeight, mapLength, i1, j1) {
					uniquePoses[pos{
						i: i1,
						j: j1,
					}] = true

				}
				i2 := two.i + (two.i - one.i)
				j2 := two.j + (two.j - one.j)
				if legal(mapHeight, mapLength, i2, j2) {
					uniquePoses[pos{
						i: i2,
						j: j2,
					}] = true

				}
			}
		}
	}
	return strconv.Itoa(len(uniquePoses))

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
