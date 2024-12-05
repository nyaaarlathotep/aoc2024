package day5

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	parts := strings.Split(input, "\n\n")
	graph := make(map[int][]int)
	for _, l := range strings.Split(parts[0], "\n") {
		dir := strings.Split(l, "|")
		one, _ := strconv.Atoi(dir[0])
		two, _ := strconv.Atoi(dir[1])
		if v, ok := graph[one]; ok {
			graph[one] = append(v, two)
		} else {
			graph[one] = []int{two}
		}
	}

	total := 0
	for _, l := range strings.Split(parts[1], "\n") {
		numsStr := strings.Split(l, ",")

		nums := make([]int, len(numsStr))
		for i := range numsStr {
			num, _ := strconv.Atoi(numsStr[i])
			nums[i] = num
		}

		pass := true
		for i := len(nums) - 1; i >= 0; i-- {
			num := nums[i]
			if nextNums, ok := graph[num]; ok {
				for j := 0; j < i; j++ {
					if slices.Contains(nextNums, nums[j]) {
						pass = false
						break
					}
				}
			}
			if !pass {
				break
			}
		}
		if pass {
			total += nums[len(nums)/2]
		}
	}
	return strconv.Itoa(total)

}

func PartTwo(input string) string {
	parts := strings.Split(input, "\n\n")
	graph := make(map[int][]int)
	for _, l := range strings.Split(parts[0], "\n") {
		dir := strings.Split(l, "|")
		one, _ := strconv.Atoi(dir[0])
		two, _ := strconv.Atoi(dir[1])
		if v, ok := graph[one]; ok {
			graph[one] = append(v, two)
		} else {
			graph[one] = []int{two}
		}
	}

	total := 0
	for _, l := range strings.Split(parts[1], "\n") {
		numsStr := strings.Split(l, ",")

		nums := make([]int, len(numsStr))
		for i := range numsStr {
			num, _ := strconv.Atoi(numsStr[i])
			nums[i] = num
		}

		pass := true
		for i := len(nums) - 1; i >= 0; i-- {
			num := nums[i]
			if nextNums, ok := graph[num]; ok {
				for j := 0; j < i; j++ {
					if slices.Contains(nextNums, nums[j]) {
						pass = false
						break
					}
				}
			}
			if !pass {
				break
			}
		}
		if !pass {
			slices.SortFunc(nums, func(a, b int) int {
				if nextNums, ok := graph[a]; ok {
					if slices.Contains(nextNums, b) {
						return -1
					}
				}
				if nextNums, ok := graph[b]; ok {
					if slices.Contains(nextNums, a) {
						return 1
					}
				}
				panic("!")
			})

			total += nums[len(nums)/2]
		}
	}
	return strconv.Itoa(total)
}
