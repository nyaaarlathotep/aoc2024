package main

import (
	"aoc2024/day17"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day17/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day17.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day17.PartTwo(string(content)))
}
