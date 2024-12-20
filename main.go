package main

import (
	"aoc2024/day20"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day20/input")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("partOne res: %s \n", day20.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day20.PartTwo(string(content)))
}
