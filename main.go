package main

import (
	"aoc2024/day16"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day16/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day16.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day16.PartTwo(string(content)))
}
