package main

import (
	"aoc2024/day6"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day6/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day6.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day6.PartTwo(string(content)))
}
