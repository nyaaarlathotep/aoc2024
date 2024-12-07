package day7

import (
	"fmt"
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
	lines := strings.Split(input, "\n")
	res := 0
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		nums := strings.Split(parts[1], " ")
		target, _ := strconv.Atoi(parts[0])
		if judge2(nums, target) {
			res += target
		}

	}
	return strconv.Itoa(res)
}

func printJudge3(nums []string, target int) bool {

	others := make([]int, len(nums))
	for i, n := range nums {
		num, _ := strconv.Atoi(n)
		others[i] = num

	}
	return printValid(target, 0, others)
}

func judge3(nums []string, target int) bool {

	others := make([]int, len(nums))
	for i, n := range nums {
		num, _ := strconv.Atoi(n)
		others[i] = num

	}
	return valid(target, 0, others)
}

// 3745369576897
func judge2(nums []string, target int) bool {

	num, _ := strconv.Atoi(nums[len(nums)-1])
	if len(nums) == 1 && num == target {
		return true
	}
	if len(nums) == 1 {
		return false
	}
	if target < num {
		return false
	}

	targetStr := strconv.Itoa(target)
	if strings.HasSuffix(targetStr, nums[len(nums)-1]) {
		newTarget, _ := strconv.Atoi(targetStr[:len(targetStr)-len(nums[len(nums)-1])])
		pass := judge2(nums[:len(nums)-1], newTarget)
		if pass {
			return pass
		}
	}
	if target%num == 0 && judge2(nums[:len(nums)-1], target/num) {
		return true
	}
	return judge2(nums[:len(nums)-1], target-num)
}

// 264184041398847
func valid(goal, temp int, others []int) bool {
	if len(others) == 0 {
		return temp == goal
	}

	if valid(goal, temp+others[0], others[1:]) {
		return true
	}
	if valid(goal, temp*others[0], others[1:]) {
		return true
	}

	temp, _ = strconv.Atoi(strconv.Itoa(temp) + strconv.Itoa(others[0]))
	return valid(goal, temp, others[1:])
}

func printValid(goal, temp int, others []int) bool {
	if len(others) == 0 {
		return temp == goal
	}

	if printValid(goal, temp+others[0], others[1:]) {
		fmt.Printf(" %d +", others[0])
		return true
	}
	if printValid(goal, temp*others[0], others[1:]) {
		fmt.Printf(" %d *", others[0])
		return true
	}

	temp, _ = strconv.Atoi(strconv.Itoa(temp) + strconv.Itoa(others[0]))
	if printValid(goal, temp, others[1:]) {
		fmt.Printf(" %d |", others[0])
		return true
	}
	return false
}
