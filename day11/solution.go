package day11

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	numStrs := strings.Split(input, " ")
	total := 0
	total = get25Times(numStrs, total)

	return strconv.Itoa(total)

}

func get25Times(numStrs []string, total int) int {
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
	return total
}

func PartTwo(input string) string {
	numStrs := strings.Split(input, " ")
	m := make(map[string]int)
	res := make([]string, 0)
	for _, v := range numStrs {
		vRes := get25TimesStrs([]string{v})
		m[v] = len(vRes)
		res = append(res, vRes...)
	}

	total := 0
	newRes := make([]string, 0)
	for _, v := range res {
		if vRes, ok := m[v]; ok {
			total = total + vRes*vRes
			continue
		}
		vRes := get25TimesStrs([]string{v})
		m[v] = len(vRes)
		newRes = append(newRes, vRes...)
	}

	for _, v := range newRes {
		if vRes, ok := m[v]; ok {
			total = total + vRes
			continue
		}
		vRes := get25TimesStrs([]string{v})
		m[v] = len(vRes)
		total += len(vRes)
	}
	return strconv.Itoa(total)

}
func get25TimesStrs(numStrs []string) []string {
	final := make([]string, 0)
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
		final = append(final, asteroids...)
	}
	return final
}
