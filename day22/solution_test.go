package day22

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
func TestPartOne1(t *testing.T) {
	fmt.Printf("partOne res: %s \n", PartOne("123"))
}
func TestPartTwo(t *testing.T) {
	content, err := os.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partTwo res: %s \n", PartTwo(string(content)))
}
