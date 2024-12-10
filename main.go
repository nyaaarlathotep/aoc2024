package main

import (
	"aoc2024/day10"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day10/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day10.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day10.PartTwo(string(content)))
}
