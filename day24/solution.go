package day24

import (
	"strconv"
	"strings"
)

type process struct {
	a, b string
	r    string
	f    func(int, int) int
}

const totalZ = 46

func PartOne(input string) string {
	parts := strings.Split(input, "\n\n")
	m := make(map[string]int)
	for _, line := range strings.Split(parts[0], "\n") {
		lar := strings.Split(line, ": ")
		l, r := lar[0], lar[1]
		var rn int
		if r == "1" {
			rn = 1
		} else {
			rn = 0
		}
		m[l] = rn
	}
	lines := strings.Split(parts[1], "\n")
	processes := make([]process, 0, len(lines))
	for _, line := range lines {
		lar := strings.Split(line, " -> ")
		l, r := lar[0], lar[1]
		lParts := strings.Split(l, " ")
		var f func(int, int) int
		if lParts[1] == "XOR" {
			f = func(i1 int, i2 int) int {
				return i1 ^ i2
			}
		} else if lParts[1] == "OR" {
			f = func(i1 int, i2 int) int {
				return i1 | i2
			}
		} else if lParts[1] == "AND" {
			f = func(i1 int, i2 int) int {
				return i1 & i2
			}
		} else {
			panic("!")
		}
		p := process{
			a: lParts[0],
			b: lParts[2],
			r: r,
			f: f,
		}
		processes = append(processes, p)
	}
	zByte := [46]int{}
	for len(processes) != 0 {
		newProcess := make([]process, 0)
		for _, p := range processes {
			va, existA := m[p.a]
			vb, existB := m[p.b]
			if !existA || !existB {
				newProcess = append(newProcess, p)
				continue
			}
			value := p.f(va, vb)
			m[p.r] = value
			if strings.HasPrefix(p.r, "z") {
				i, _ := strconv.Atoi(p.r[1:])
				zByte[i] = value
			}
		}
		processes = newProcess
	}

	total := 0
	for i := 0; i < totalZ; i++ {
		b := zByte[i]
		total += b << i
	}
	return strconv.Itoa(total)

}

func PartTwo(input string) string {
	return ""

}
