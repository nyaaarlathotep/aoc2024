package day24

import (
	"fmt"
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
	testM := make(map[string]int)
	for i := 0; i <= 45; i++ {
		var x, y, z string
		if i/10 > 0 {
			x = "x" + strconv.Itoa(i)
			y = "y" + strconv.Itoa(i)
			z = "z" + strconv.Itoa(i)
		} else {
			y = "y0" + strconv.Itoa(i)
			x = "x0" + strconv.Itoa(i)
			z = "z0" + strconv.Itoa(i)
		}
		testM[x] = 1
		testM[y] = 1
		fmt.Printf("\n%v -> ", i)

		for len(processes) != 0 {
			newProcess := make([]process, 0)
			for _, p := range processes {
				va, existA := testM[p.a]
				vb, existB := testM[p.b]
				if !existA || !existB {
					newProcess = append(newProcess, p)
					continue
				}
				fmt.Printf("%v,%v -> %v ", p.a, p.b, p.r)
				value := p.f(va, vb)
				testM[p.r] = value
			}
			if len(processes) == len(newProcess) {
				break
			}
			processes = newProcess
		}

		if v, ok := testM[z]; ok {
			if v == 0 {
				fmt.Println("!")
			}
		} else {
			fmt.Println("!")
		}

		// ((x and y) xor ) xor (x xor y)
		// z12,jpj
		// chv,vvw

	}

	return swapAndJoinWires()

}

func getIndex(s string) int {
	i, _ := strconv.Atoi(s[1:])
	return i
}
