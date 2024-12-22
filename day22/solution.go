package day22

import (
	"strconv"
	"strings"
)

const COUNT = 2000

func PartOne(input string) string {
	total := 0
	for _, l := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(l)
		final := getNth(num, COUNT)
		total += final
	}
	return strconv.Itoa(total)

}

func getNth(n, count int) int {
	for range count {
		mix := n * 64
		step1 := mixAndPurge(n, mix)
		step2 := mixAndPurge(step1, step1/32)
		step3 := mixAndPurge(step2, step2*2048)
		n = step3
	}
	return n
}

func mixAndPurge(n int, mix int) int {
	mixed := n ^ mix
	step1 := mixed % 16777216
	return step1
}

func PartTwo(input string) string {
	return ""

}
