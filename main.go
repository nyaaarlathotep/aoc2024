package main

import (
	"aoc2024/day7"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day7/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day7.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day7.PartTwo(string(content)))
}
