package main

import (
	"aoc2024/day12"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day12/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day12.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day12.PartTwo(string(content)))
}
