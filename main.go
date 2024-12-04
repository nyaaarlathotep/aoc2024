package main

import (
	"aoc2024/day4"
	"fmt"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	content, err := os.ReadFile("./day4/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", day4.PartOne(string(content)))
	fmt.Printf("partTwo res: %s \n", day4.PartTwo(string(content)))
}
