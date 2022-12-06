package day_seven

import (
	"bufio"
	"fmt"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_seven/test.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		fmt.Println(line)
	}
}