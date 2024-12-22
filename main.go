package main

import (
	"aoc2024/day22"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./day22/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day22.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day22.PartTwo(string(content)))
}
