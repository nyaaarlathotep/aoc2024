package day1

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func PartOne(input string) string {

	split := strings.Split(input, "\n")
	one := make([]int, len(split))
	two := make([]int, len(split))
	for i, line := range split {
		re := regexp.MustCompile(`(\d+)\s+(\d+)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			num1, _ := strconv.Atoi(matches[1])
			one[i] = num1
			num2, _ := strconv.Atoi(matches[2])
			two[i] = num2
		}
	}

	sort.Ints(one)
	sort.Ints(two)
	total := 0
	for i := range one {
		if one[i] > two[i] {
			total += one[i] - two[i]
		} else {
			total += two[i] - one[i]
		}
	}
	return strconv.Itoa(total)
}
