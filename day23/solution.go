package day23

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	m := make(map[string]map[string]bool)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		l, r := parts[0], parts[1]
		if m[l] == nil {
			m[l] = make(map[string]bool)
		}
		if m[r] == nil {
			m[r] = make(map[string]bool)
		}
		m[l][r] = true
		m[r][l] = true
	}
	all := allNSets(m, 2)
	return strconv.Itoa(len(all))

}

func allNSets(m map[string]map[string]bool, n int) []string {
	finalRes := make(map[string]bool)
	for k := range m {
		if !strings.HasPrefix(k, "t") {
			continue
		}
		exist := make(map[string]bool)
		exist[k] = true
		res := innerNSet(m, exist, k, n)
		for _, v := range res {
			v = append(v, k)
			slices.Sort(v)
			finalRes[strings.Join(v, ",")] = true
		}
	}
	nets := make([]string, 0, len(finalRes))
	for k := range finalRes {
		nets = append(nets, k)
	}
	return nets
}

func innerNSet(m map[string]map[string]bool, exist map[string]bool, now string, remainN int) [][]string {
	if remainN == 0 {
		return [][]string{{}}
	}
	neighbors := m[now]
	res := make([][]string, 0)
	for neighbor := range neighbors {
		if exist[neighbor] {
			continue
		}
		connect := true
		nNeighbors := m[neighbor]
		for existNeighbor, v := range exist {
			if !v {
				continue
			}
			if !nNeighbors[existNeighbor] {
				connect = false
				break
			}
		}
		if !connect {
			continue
		}
		exist[neighbor] = true
		remainRes := innerNSet(m, exist, neighbor, remainN-1)
		exist[neighbor] = false
		for _, e := range remainRes {
			res = append(res, append(e, neighbor))
		}
	}
	return res
}

func PartTwo(input string) string {
	return ""

}
