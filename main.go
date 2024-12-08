package main

import (
	"aoc2024/day8"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day8/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day8.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day8.PartTwo(string(content)))
}
