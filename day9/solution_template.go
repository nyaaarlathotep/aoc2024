package day9

import (
	"fmt"
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

func PartTwo(input string) string {
	res := make([]space, len(input))
	spaceType := true
	fileId := 0
	for i, c := range input {
		if spaceType {
			res[i] = space{
				id:        fileId,
				length:    int(c - '0'),
				spaceType: file,
			}
			fileId++
			spaceType = false
		} else {
			res[i] = space{
				id:        -1,
				length:    int(c - '0'),
				spaceType: empty,
			}
			spaceType = true
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		if res[i].spaceType == file {
			for j := 0; j < i; j++ {
				if res[j].spaceType == empty && res[j].length >= res[i].length {
					res[i].spaceType = empty
					newFile := space{
						id:        res[i].id,
						length:    res[i].length,
						spaceType: file,
						children:  nil,
					}
					res[i].id = -1
					if res[j].length == res[i].length {
						res[j] = newFile
						//printRes(&res)
						break
					}
					res[j].length = res[j].length - res[i].length
					res = append(res[:j+1], res[j:]...)
					res[j] = newFile
					//printRes(&res)
					break
				}
			}
		}

	}

	total := 0
	index := 0
	for _, v := range res {
		for range v.length {
			if v.spaceType == file {
				total = total + index*v.id
			}
			index++
		}
	}
	return strconv.Itoa(total)
}

func printRes(resa *[]space) {
	res := *resa
	for _, v := range res {
		for range v.length {
			fmt.Printf("%d ", v.id)
		}
	}
	fmt.Println()
}

func write(s *[]int, i int, length int, fileId int) {
	for j := i; j < i+length; j++ {
		(*s)[j] = fileId
	}
}
func findSpace(s *[]int, i int, length int) bool {
	for j := i; j < i+length; j++ {
		if (*s)[j] != -1 {
			return false
		}
	}
	return true
}
