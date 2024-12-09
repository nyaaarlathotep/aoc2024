package main

import (
	"aoc2024/day9"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day9/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day9.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day9.PartTwo(string(content)))
}
