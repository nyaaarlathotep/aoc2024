package main

import (
	"aoc2024/day13"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day13/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day13.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day13.PartTwo(string(content)))
}
