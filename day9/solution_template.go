package day9

import (
	"slices"
	"strconv"
)

const file = 1
const empty = 2

type space struct {
	id        int
	length    int
	spaceType int
	children  []space
}

func PartOne(input string) string {
	res := make([]int, 0)
	fileId := 0
	spaceType := true

	for _, c := range input {
		length := int(c - '0')
		if spaceType {
			slices.Grow(res, length)
			for j := 0; j < length; j++ {
				res = append(res, fileId)
			}
			fileId++
			spaceType = false
		} else {
			spaceType = true
			slices.Grow(res, length)
			for j := 0; j < length; j++ {
				res = append(res, -1)
			}
		}
	}
	low, high := 0, len(res)-1
	for low < high {
		for res[low] != -1 {
			low++
			if low >= high {
				break
			}
		}
		for res[high] == -1 {
			high--
			if low >= high {
				break
			}
		}
		if low >= high {
			break
		}
		res[low], res[high] = res[high], res[low]
	}
	total := 0

	for i := range res {
		if res[i] == -1 {
			break
		}
		total += i * res[i]
	}

	return strconv.Itoa(total)
}

//func PartOne(input string) string {
//
//	res := make([]space, 0)
//
//	spaceType := true
//	fileId := 0
//	for i, c := range input {
//		if spaceType {
//			res[i] = space{
//				id:        fileId,
//				length:    int(c - '0'),
//				spaceType: file,
//			}
//			fileId++
//			spaceType = false
//		} else {
//			res[i] = space{
//				length:    int(c - '0'),
//				spaceType: empty,
//			}
//			spaceType = true
//		}
//	}
//
//	low, high := 0, len(res)-1
//	for low < high {
//		for res[low].spaceType != empty {
//			low++
//			if low >= high {
//				break
//			}
//		}
//		for res[high].spaceType != file {
//			high--
//			if low >= high {
//				break
//			}
//		}
//		if res[low].length == res[high].length {
//			res[low], res[high] = res[high], res[low]
//		} else if res[low].length < res[high].length {
//			if res[low].children != nil {
//
//			} else {
//				res[low].id = res[high].id
//				res[low].spaceType = res[high].spaceType
//				res[high].length = res[high].length - res[low].length
//			}
//		} else {
//			res[low].children = []space{res[high]}
//			res[low].length = res[low].length - res[high].length
//		}
//	}
//
//	for _, v := range res {
//
//	}

//return ""
//
//}

func PartTwo(input string) string {
	return ""
}
