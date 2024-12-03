package day3

import (
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
	//content, err := os.ReadFile("./test")
	//if err != nil {
	//	panic(err)
	//}
	content := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	fmt.Printf("partTwo res: %s \n", PartTwo(string(content)))
}
