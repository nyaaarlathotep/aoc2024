package main

import (
	"aoc2024/day18"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day18/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day18.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day18.PartTwo(string(content)))
}
