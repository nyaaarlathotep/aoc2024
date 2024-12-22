package day22

import (
	"fmt"
	"math"
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

type seq struct {
	a, b, c, d int
}

// PartTwo 1454 is wrong, why?
// damn it, it would wait until the first time it sees that sequence and then immediately sell
func PartTwo(input string) string {
	split := strings.Split(input, "\n")
	seqMaps := make([]map[seq]int, 0, len(split))
	secrets := make([]int, 0)
	for _, l := range split {
		num, _ := strconv.Atoi(l)
		m := make(map[seq]int)
		getNth2(num, COUNT, m)
		seqMaps = append(seqMaps, m)

		secrets = append(secrets, num)
	}

	allKinds := make(map[seq]bool)
	for _, m := range seqMaps {
		for k := range m {
			allKinds[k] = true
		}
	}

	maxBananas := math.MinInt
	var maxKind seq
	for k := range allKinds {

		thisKindCount := 0
		for _, m := range seqMaps {
			thisKindCount += m[k]
		}
		if thisKindCount > maxBananas {
			maxBananas = thisKindCount
			maxKind = k
		}
	}
	fmt.Println(maxKind)
	return strconv.Itoa(maxBananas)
}

func getNth2(n, count int, seqMap map[seq]int) int {
	tempS := make([]int, 0, 5)
	fistLastDigit := n % 10
	tempS = append(tempS, fistLastDigit)
	for i := range count {
		mix := n * 64
		step1 := mixAndPurge(n, mix)
		step2 := mixAndPurge(step1, step1/32)
		step3 := mixAndPurge(step2, step2*2048)
		n = step3

		lastDigit := n % 10
		tempS = append(tempS, lastDigit)
		if i > 2 {
			s := seq{
				a: tempS[1] - tempS[0],
				b: tempS[2] - tempS[1],
				c: tempS[3] - tempS[2],
				d: tempS[4] - tempS[3],
			}
			tempS = tempS[1:]
			if _, ok := seqMap[s]; ok {
			} else {
				seqMap[s] = lastDigit
			}
		}
	}
	return n
}
