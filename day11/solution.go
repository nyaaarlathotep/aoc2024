package day11

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	numStrs := strings.Split(input, " ")
	total := 0
	for _, num := range numStrs {
		asteroids := []string{num}
		for range 25 {
			next := make([]string, 0)
			slices.Grow(next, len(asteroids))
			for _, a := range asteroids {
				if a == "0" {
					next = append(next, "1")
				} else if len(a)%2 == 0 {
					left, _ := strconv.Atoi(a[:len(a)/2])
					right, _ := strconv.Atoi(a[len(a)/2:])
					next = append(next, strconv.Itoa(left))
					next = append(next, strconv.Itoa(right))
				} else {
					v, _ := strconv.Atoi(a)
					next = append(next, strconv.Itoa(v*2024))
				}
			}
			asteroids = next
		}
		total = total + len(asteroids)
	}

	return strconv.Itoa(total)

}

func PartTwo(input string) string {
	return ""

}
