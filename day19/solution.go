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
	return ""

}
