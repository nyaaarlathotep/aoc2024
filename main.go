package main

import (
	"aoc2024/day24"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day24/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day24.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day24.PartTwo(string(content)))
}
