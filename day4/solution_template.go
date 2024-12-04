package day4

import (
	"strconv"
	"strings"
)

func PartOne(input string) string {
	lines := strings.Split(input, "\n")
	byteMap := make([][]byte, len(lines))
	for i, l := range lines {
		byteMap[i] = make([]byte, len(l))
		for j, b := range []byte(l) {
			byteMap[i][j] = b
		}
	}
	total := 0
	for i := range byteMap {
		for j := range byteMap[0] {
			if byteMap[i][j] == 'X' {
				allWords := allDirBytes(byteMap, 4, i, j)
				count := 0
				for _, word := range allWords {
					if equalXmas(word) {
						count++
					}
				}

				total += count
			}
		}
	}
	return strconv.Itoa(total)
}

func equalXmas(b []byte) bool {
	return b[0] == 'X' && b[1] == 'M' && b[2] == 'A' && b[3] == 'S'
}

func allDirBytes(m [][]byte, targetLen int, i, j int) [][]byte {
	res := make([][]byte, 0)
	if i-targetLen+1 >= 0 {
		up := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			up[n] = m[ti][tj]
			ti--
		}
		res = append(res, up)
	}
	if j-targetLen+1 >= 0 {
		left := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			left[n] = m[ti][tj]
			tj--
		}
		res = append(res, left)
	}
	if j-targetLen+1 >= 0 && i-targetLen+1 >= 0 {
		upLeft := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			upLeft[n] = m[ti][tj]
			ti--
			tj--
		}
		res = append(res, upLeft)
	}

	if i+targetLen <= len(m) {
		down := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			down[n] = m[ti][tj]
			ti++
		}
		res = append(res, down)
	}
	if j+targetLen <= len(m) {
		right := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			right[n] = m[ti][tj]
			tj++
		}
		res = append(res, right)
	}

	if j+targetLen <= len(m) && i-targetLen+1 >= 0 {
		upRight := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			upRight[n] = m[ti][tj]
			tj++
			ti--
		}
		res = append(res, upRight)
	}

	if j+targetLen <= len(m) && i+targetLen <= len(m) {
		downRight := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			downRight[n] = m[ti][tj]
			tj++
			ti++
		}
		res = append(res, downRight)
	}

	if i+targetLen <= len(m) && j-targetLen+1 >= 0 {
		downLeft := make([]byte, targetLen)
		ti, tj := i, j
		for n := 0; n < targetLen; n++ {
			downLeft[n] = m[ti][tj]
			ti++
			tj--
		}
		res = append(res, downLeft)
	}

	return res
}

func PartTwo(input string) string {
	lines := strings.Split(input, "\n")
	byteMap := make([][]byte, len(lines))
	for i, l := range lines {
		byteMap[i] = make([]byte, len(l))
		for j, b := range []byte(l) {
			byteMap[i][j] = b
		}
	}
	total := 0
	for i := 1; i < len(byteMap)-1; i++ {
		for j := 1; j < len(byteMap[0])-1; j++ {
			if byteMap[i][j] == 'A' {
				word1 := [2]byte{byteMap[i-1][j-1], byteMap[i+1][j+1]}
				word2 := [2]byte{byteMap[i-1][j+1], byteMap[i+1][j-1]}
				if digPass(word1) && digPass(word2) {
					total++
				}

			}
		}
	}
	return strconv.Itoa(total)
}

func digPass(bytes [2]byte) bool {
	return (bytes[0] == 'M' && bytes[1] == 'S') || (bytes[1] == 'M' && bytes[0] == 'S')
}
