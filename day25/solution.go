package day25

import (
	"strconv"
	"strings"
)

type shape struct {
	a, b, c, d, e int
}

func PartOne(input string) string {

	parts := strings.Split(input, "\n\n")
	keys := make([][5]int, 0)
	locks := make([][5]int, 0)
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		lens := [5]int{}
		if lines[0][0] == '#' {
			for j := range lines[0] {
				length := 0
				for i := range lines {
					if lines[i][j] != '#' {
						break
					}
					length++

				}
				lens[j] = length - 1
			}
			locks = append(locks, lens)
		} else {
			for j := range lines[0] {
				length := 0
				for i := range lines {
					if lines[i][j] != '.' {
						break
					}
					length++
				}
				lens[j] = length - 1
			}
			keys = append(keys, lens)
		}
	}

	total := 0
	for _, lock := range locks {
		for _, key := range keys {
			pass := true
			for i := range lock {
				if key[i] < lock[i] {
					pass = false
					break
				}
			}
			if pass {
				total++
			}
		}
	}
	return strconv.Itoa(total)
}

func PartTwo(input string) string {
	return ""

}
