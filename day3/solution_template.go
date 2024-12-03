package day3

import (
	"regexp"
	"strconv"
)

func PartOne(input string) string {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	res := 0
	for _, match := range matches {
		one, _ := strconv.Atoi(match[1])
		two, _ := strconv.Atoi(match[2])
		res = res + one*two
	}
	return strconv.Itoa(res)
}

func PartTwo(input string) string {
	return ""

}
