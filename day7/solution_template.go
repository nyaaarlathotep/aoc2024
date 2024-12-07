package day7

import (
	"strconv"
	"strings"
)

func PartOne(input string) string {
	lines := strings.Split(input, "\n")
	res := 0
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		nums := strings.Split(parts[1], " ")
		target, _ := strconv.Atoi(parts[0])
		if judge(nums, target) {
			res += target
		}

	}
	return strconv.Itoa(res)
}

func judge(nums []string, target int) bool {

	num, _ := strconv.Atoi(nums[len(nums)-1])
	if num == target {
		return true
	}
	if len(nums) == 1 {
		return false
	}
	if target < num {
		return false
	}
	if target%num == 0 {
		return judge(nums[:len(nums)-1], target/num) || judge(nums[:len(nums)-1], target-num)
	}
	return judge(nums[:len(nums)-1], target-num)
}

func PartTwo(input string) string {
	return ""

}
