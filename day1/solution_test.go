package day1

import (
	"fmt"
	"os"
	"testing"
)

func TestPartTwo(t *testing.T) {
	content, err := os.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("partOne res: %s \n", PartTwo(string(content)))
}
