package day_five

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_five/data.txt"
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
			continue
		}
		howMany, from, to := parseMove(line)
		columns = moveMulti(howMany, from, to, columns)	
	}
	topRow := make([]string, numCols)
	for i, v := range columns {
		topRow[i], _ = pop(v)
	} 
	fmt.Printf("Top Row: %+v\n", topRow)
}

func parseMove(command string) (int, int, int) {
	regex := regexp.MustCompile(`\d+`)
	nums := regex.FindAllString(command, -1)
	if nums == nil {
		fmt.Println("Uh oh, something's wrong...")
		return 0,0,0
	}
	howmany, _ := strconv.Atoi(nums[0])
	from, _ := strconv.Atoi(nums[1])
	to, _ := strconv.Atoi(nums[2])
	return howmany, from -1, to -1
}

func moveMulti(howMany int, from int, to int, columns [][]string) [][]string {
	for i := 0; i < howMany; i++ {
		columns = moveSingle(from, to, columns)
	}
	return columns
}

func moveSingle(from int, to int, columns [][]string) [][]string {
	box, newFrom := pop(columns[from])
	newTo := append(columns[to], box)
	columns[from] = newFrom
	columns[to] = newTo
	return columns
}

func pop(col []string) (string, []string) {
	newLen := len(col) -1
	value := col[newLen]
	newCol := col[:newLen]
	return value, newCol
}

func initializeStacks(stackStrings []string, numCols int) [][]string {
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