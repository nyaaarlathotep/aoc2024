package main

import (
	"aoc2024/day21"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day21/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day21.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day21.PartTwo(string(content)))
}
