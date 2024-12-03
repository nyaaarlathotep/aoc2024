package main

import (
	"aoc2024/day3"
	"fmt"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	content, err := os.ReadFile("./day3/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day3.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day3.PartTwo(string(content)))
}
