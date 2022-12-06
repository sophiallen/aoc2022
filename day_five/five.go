package day_five

import (
	"bufio"
	"fmt"
	"strings"
	// "regexp"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_five/test.txt"
}

// A = 65
// Z == 90
// a == 97
// z == 122

func (s Solution) Solve(scanner *bufio.Scanner) {
	// stacks := [][]string{}
	stackStrings := []string{}
	numCols := 0
	initialized := false
	columns := [][]string{}

	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		if !initialized {
			isBoxLine := strings.Contains(line, "[")
			if isBoxLine {
				stackStrings = append(stackStrings, line)
				continue
			}
			if len(line) > 0 {
				// it's the column numbers line
				cols := strings.ReplaceAll(line, " ", "")
				numCols = len(cols)
				continue
			}
			columns = initializeStacks(stackStrings, numCols)
			initialized = true
		}
		if initialized {
			fmt.Printf("Columns are %+v\n", columns)
			break
		}

		fmt.Println(line)
	}
}

func initializeStacks(stackStrings []string, numCols int) [][]string {
	// boxRegex := regexp.MustCompile(`\[.{1}\`)
	rows := make([][]string, len(stackStrings))
	for i, txt := range stackStrings {
		row := lineToRow(txt, numCols)
		fmt.Println(row)
		rows[i] = row
	}
	return rowsToColumns(rows, numCols)
}

func lineToRow(line string, numCols int) []string {
	row := make([]string, numCols)
	for i, r := range line {
		if r >= 65 && r <= 90 {
			// found box
			col := i / 4
			row[col] = string(r)
		}
	}
	return row
}

func rowsToColumns(rows [][]string, numCols int) [][]string {
	stacks := [][]string{}
	// for each column
	for col := 0; col < numCols; col++ {
		column := []string{}
		// check rows from bottomw to top (reverse order)
		for r := len(rows) - 1; r >= 0; r-- {
			row := rows[r]
			if len(row[col]) > 0 {
				// if letter exists, add to top of column
				column = append(column, row[col])
			}
		}
		stacks = append(stacks, column)
	}
	return stacks
}
