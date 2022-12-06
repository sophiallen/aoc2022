package main

import (
	"aoc2022/day_five"
	"aoc2022/day_four"
	"aoc2022/day_one"
	"aoc2022/day_three"
	"aoc2022/day_two"
	"bufio"
	"fmt"
	"os"
)

type Solvable interface {
	Solve(scanner *bufio.Scanner)
	GetDataPath() string
}

var solutions = map[int]Solvable{
	1: day_one.Solution{},
	2: day_two.Solution{},
	3: day_three.Solution{},
	4: day_four.Solution{},
	5: day_five.Solution{},
}

func main() {
	const day = 5
	solution := solutions[day]
	fileScanner, reader := openFile(solution.GetDataPath())
	solution.Solve(fileScanner)
	reader.Close()
}

func openFile(file string) (*bufio.Scanner, *os.File) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readFile
}
