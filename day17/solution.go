package day17

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	a, b, c := 0, 0, 0
	parts := strings.Split(input, "\n\n")
	for i, l := range strings.Split(parts[0], "\n") {
		lr := strings.Split(l, ": ")
		n, _ := strconv.Atoi(lr[1])
		if i == 0 {
			a = n
		} else if i == 1 {
			b = n
		} else if i == 2 {
			c = n
		}
	}
	rawPrograms := strings.Split(parts[1], ": ")[1]
	programs := make([]int, 0)
	for _, r := range strings.Split(rawPrograms, ",") {
		n, _ := strconv.Atoi(r)
		programs = append(programs, n)
	}

	res := getRes(a, b, c, programs)
	strRes := make([]string, 0, len(res))
	for _, s := range res {
		strRes = append(strRes, strconv.Itoa(s))
	}
	return strings.Join(strRes, ",")
}

func getRes(a int, b int, c int, programs []int) []int {
	comboValue := func(n int) int {
		if n <= 3 {
			return n
		}
		if n == 4 {
			return a
		}
		if n == 5 {
			return b
		}
		if n == 6 {
			return c
		}
		panic(n)
	}
	adv := func(n int) {
		actualN := comboValue(n)
		a = a / twoPow(actualN)
	}
	bxl := func(n int) {
		b = n ^ b
	}
	bst := func(n int) {
		actualN := comboValue(n)
		b = actualN % 8
	}
	jnz := func(n int) int {
		if a == 0 {
			return -1
		}
		return n
	}
	bxc := func(n int) {
		b = b ^ c
	}
	out := func(n int) int {
		actualN := comboValue(n)
		return actualN % 8
	}
	bdv := func(n int) {
		actualN := comboValue(n)
		b = a / twoPow(actualN)
	}
	cdv := func(n int) {
		actualN := comboValue(n)
		c = a / twoPow(actualN)
	}
	m := make(map[int]func(n int))
	m[0] = adv
	m[1] = bxl
	m[2] = bst
	//m[3] = jnz
	m[4] = bxc
	//m[5] = out
	m[6] = bdv
	m[7] = cdv
	res := make([]int, 0)
	for i := 0; i < len(programs); {
		operator := programs[i]
		operand := programs[i+1]
		if operator == 3 {
			jump := jnz(operand)
			if jump != -1 {
				i = jump
			} else {
				i = i + 2
			}
			continue
		}
		if operator == 5 {
			res = append(res, out(operand))
		} else {
			m[operator](operand)
		}
		i = i + 2
	}
	return res
}

func twoPow(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * twoPow(n-1)
}

func PartTwo(input string) string {
	a, b, c := 0, 0, 0
	parts := strings.Split(input, "\n\n")
	for i, l := range strings.Split(parts[0], "\n") {
		lr := strings.Split(l, ": ")
		n, _ := strconv.Atoi(lr[1])
		if i == 0 {
			a = n
		} else if i == 1 {
			b = n
		} else if i == 2 {
			c = n
		}
	}
	rawPrograms := strings.Split(parts[1], ": ")[1]
	programs := make([]int, 0)
	for _, r := range strings.Split(rawPrograms, ",") {
		n, _ := strconv.Atoi(r)
		programs = append(programs, n)
	}
	comboValue := func(n int) int {
		if n <= 3 {
			return n
		}
		if n == 4 {
			return a
		}
		if n == 5 {
			return b
		}
		if n == 6 {
			return c
		}
		panic(n)
	}
	adv := func(n int) {
		actualN := comboValue(n)
		a = a / twoPow(actualN)
	}
	bxl := func(n int) {
		b = n ^ b
	}
	bst := func(n int) {
		actualN := comboValue(n)
		b = actualN % 8
	}
	jnz := func(n int) int {
		if a == 0 {
			return -1
		}
		return n
	}
	bxc := func(n int) {
		b = b ^ c
	}
	out := func(n int) int {
		actualN := comboValue(n)
		return actualN % 8
	}
	bdv := func(n int) {
		actualN := comboValue(n)
		b = a / twoPow(actualN)
	}
	cdv := func(n int) {
		actualN := comboValue(n)
		c = a / twoPow(actualN)
	}
	m := make(map[int]func(n int))
	m[0] = adv
	m[1] = bxl
	m[2] = bst
	//m[3] = jnz
	m[4] = bxc
	//m[5] = out
	m[6] = bdv
	m[7] = cdv
	//lastSuccess := 821010000
	lastSuccess := 2
	powTimes := 1
	res := getRes(a, 0, 0, programs)
	a = 0
	for n := len(programs) - 1; n >= 0; n-- {
		a = a << 3
		for !slices.Equal(getRes(a, 0, 0, programs), programs) {
			a++
		}
	}

	fmt.Println(lastSuccess)
	fmt.Println(res)

	lastSuccess = lastSuccess * 8
	//lastSuccess = int(math.Pow(8.0, float64(powTimes)))
	powTimes++
	return ""
}
