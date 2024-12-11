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
	return get75TimesStrs(numStrs)

}
func get75TimesStrs(numStrs []string) string {
	asteroids := make(map[string]int)
	for _, a := range numStrs {
		if v, ok := asteroids[a]; ok {
			asteroids[a] = v + 1
		} else {
			asteroids[a] = 1
		}
	}
	for range 75 {
		next := make(map[string]int)
		for k, v := range asteroids {
			if k == "0" {
				nextKey := "1"
				mapAddOrSet(next, nextKey, v)
			} else if len(k)%2 == 0 {
				left, _ := strconv.Atoi(k[:len(k)/2])
				right, _ := strconv.Atoi(k[len(k)/2:])

				mapAddOrSet(next, strconv.Itoa(left), v)
				mapAddOrSet(next, strconv.Itoa(right), v)
			} else {
				num, _ := strconv.Atoi(k)
				mapAddOrSet(next, strconv.Itoa(num*2024), v)
			}
		}
		asteroids = next
	}
	total := 0
	for _, v := range asteroids {
		total += v

	}
	return strconv.Itoa(total)
}

func mapAddOrSet(next map[string]int, nextKey string, v int) {
	if nextV, ok := next[nextKey]; ok {
		next[nextKey] = nextV + v
	} else {
		next[nextKey] = v
	}
}
