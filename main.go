package main

import (
	"aoc2024/day23"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day23/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day23.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day23.PartTwo(string(content)))
}
