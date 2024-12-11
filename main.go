package main

import (
	"aoc2024/day11"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day11/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day11.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day11.PartTwo(string(content)))
}
