package day2

import (
	"math"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, " ")
		safe := true
		one, _ := strconv.Atoi(nums[0])
		two, _ := strconv.Atoi(nums[1])
		ascend := two > one
		for i := 0; i < len(nums)-1; i++ {
			one, _ := strconv.Atoi(nums[i])
			two, _ := strconv.Atoi(nums[i+1])
			if two > one && !ascend {
				safe = false
				break

			} else if one > two && ascend {
				safe = false
				break
			}
			if !(math.Abs(float64(one-two)) <= 3 && math.Abs(float64(one-two)) >= 1) {
				safe = false
				break
			}
		}
		if safe {
			count++
		}
	}
	return strconv.Itoa(count)

}

func PartTwo(input string) string {
	return ""
}
