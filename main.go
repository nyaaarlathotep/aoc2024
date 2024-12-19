package main

import (
	"aoc2024/day19"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day19/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day19.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day19.PartTwo(string(content)))
}
