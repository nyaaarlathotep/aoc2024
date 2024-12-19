package day19

import (
	"strconv"
	"strings"
)

func PartOne(input string) string {
	parts := strings.Split(input, "\n\n")
	towels := make(map[string]bool)
	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel] = true
	}
	count := 0
	for _, t := range strings.Split(parts[1], "\n") {
		if exist(t, &towels) {
			count++
		}
	}
	return strconv.Itoa(count)

}

func exist(towel string, towels *map[string]bool) bool {
	if len(towel) == 0 {
		return true
	}
	for k := range *towels {
		if strings.HasPrefix(towel, k) {
			remainExist := exist(towel[len(k):], towels)
			if remainExist {
				return true
			}
		}
	}
	return false
}

func PartTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	towels := make(map[string]bool)
	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel] = true
	}
	count := 0
	for _, t := range strings.Split(parts[1], "\n") {
		count = count + poss(t, towels, make(map[string]int))
	}
	return strconv.Itoa(count)

}
func poss(towel string, towels map[string]bool, cache map[string]int) int {
	if len(towel) == 0 {
		return 1
	}
	if v, ok := cache[towel]; ok {
		return v
	}
	count := 0
	for k := range towels {
		if strings.HasPrefix(towel, k) {
			remainCount := poss(towel[len(k):], towels, cache)
			count = count + remainCount
		}
	}
	cache[towel] = count
	return count
}
