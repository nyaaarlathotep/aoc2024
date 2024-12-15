package main

import (
	"aoc2024/day15"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day15/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day15.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day15.PartTwo(string(content)))
}
