package main

import (
	"aoc2024/day5"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day5/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day5.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day5.PartTwo(string(content)))
}
