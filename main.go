package main

import (
	"aoc2024/day25"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day25/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day25.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day25.PartTwo(string(content)))
}
