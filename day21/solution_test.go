package day21

import (
	"aoc2024/runeMap"
	"fmt"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	content, err := os.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", PartOne(string(content)))
}

func TestPartTwo(t *testing.T) {
	content, err := os.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partTwo res: %s \n", PartTwo(string(content)))
}
func TestGetALlMoves(t *testing.T) {
	avoid := runeMap.Pos{
		I: 3,
		J: 0,
	}
	moves := getAllPossibleMove(avoid, letterMap['1'], letterMap['9'])
	fmt.Printf("moves res: %v \n", moves)
}
