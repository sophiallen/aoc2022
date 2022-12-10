package day_eight

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_eight/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	grid := [][]int{}
	visibleTrees := []int{}

	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		fmt.Println(line)
		rowStr := strings.Split(line, "")
		rowInt := make([]int, len(rowStr))
		for i, v := range rowStr {
			rowInt[i], _ = strconv.Atoi(v)
		}
		grid = append(grid, rowInt)
	}
	maxScenicScore := 0
	for r := 0; r < len(grid); r++ {
		row := grid[r]
		for c := 0; c < len(row); c++ {
			canBeSeen, scenicScore := isVisible(r, c, grid)
			if canBeSeen {
				visibleTrees = append(visibleTrees, grid[r][c])
			}
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(visibleTrees)
	fmt.Printf("Total visible is %+v", len(visibleTrees))
	fmt.Printf("Max scenic score is %+v", maxScenicScore)
}

func isVisibleUp(row int, col int, grid [][]int) (bool, int) {
	if row == 0 {
		return true, 0
	}
	visibleCount := 0
	height := grid[row][col]
	for r := row - 1; r >= 0; r-- {
		if grid[r][col] >= height {
			return false, visibleCount + 1
		}
		visibleCount += 1
	}
	return true, visibleCount
}

func isVisibleDown(row int, col int, grid [][]int) (bool, int) {
	if row == len(grid[0])-1 {
		return true, 0
	}
	distance := 0
	height := grid[row][col]
	for r := row + 1; r < len(grid); r++ {
		if grid[r][col] >= height {
			return false, distance + 1
		}
		distance += 1
	}
	return true, distance
}

func isVisibleLeft(row int, col int, grid [][]int) (bool, int) {
	if col == 0 {
		return true, 0
	}
	distance := 0
	height := grid[row][col]
	r := grid[row]
	for i := col - 1; i >= 0; i-- {
		if r[i] >= height {
			return false, distance + 1
		}
		distance += 1
	}

	return true, distance
}

func isVisibleRight(row int, col int, grid [][]int) (bool, int) {
	if col == len(grid[row])-1 {
		return true, 0
	}
	height := grid[row][col]
	distance := 0
	r := grid[row]
	for i := col + 1; i < len(r); i++ {
		distance += 1
		if r[i] >= height {
			return false, distance
		}
	}
	return true, distance
}

func isVisible(row int, col int, grid [][]int) (bool, int) {
	top, dt := isVisibleUp(row, col, grid)
	left, dl := isVisibleLeft(row, col, grid)
	right, dr := isVisibleRight(row, col, grid)
	bottom, db := isVisibleDown(row, col, grid)

	canBeSeen := top || left || right || bottom
	scenicScore := dt * dl * dr * db
	// fmt.Printf("(%+v, %+v)'s scenic score is %+v: %+v %+v %+v %+v\n", row, col, scenicScore, dt, dl, dr, db)
	return canBeSeen, scenicScore
}
