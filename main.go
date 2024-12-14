package main

import (
	"aoc2024/day14"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day14/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day14.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day14.PartTwo(string(content)))
}
