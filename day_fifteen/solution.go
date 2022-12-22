package day_fifteen

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_fifteen/test.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	sensorPositions := []Position{}
	beaconPositions := []Position{}
	var min, max Position
	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		// fmt.Println(line)
		positions := strings.Split(line, ":")
		sensorPos := lineToPosition(positions[0])
		beaconPos := lineToPosition(positions[1])

		sensorPos.SetNearest(beaconPos)
		sensorPositions = append(sensorPositions, sensorPos)
		beaconPositions = append(beaconPositions, beaconPos)
		min, max = updateMinMax(sensorPos, min, max)
		min, max = updateMinMax(beaconPos, min, max)
	}
	for _, p := range sensorPositions {
		fmt.Println(p)
	}

	fmt.Println("-------")
	// for _, b := range beaconPositions {
	// 	fmt.Println(b)
	// }
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func lineToPosition(line string) Position {
	parts := strings.Split(line, ",")
	xParts := strings.Split(parts[0], "x=")
	yParts := strings.Split(parts[1], "y=")
	x, err := strconv.Atoi(xParts[1])
	if err != nil {
		fmt.Printf("oh noes")
	}
	y, err := strconv.Atoi(yParts[1])
	if err != nil {
		fmt.Printf("oops")
	}
	return Position{
		Row: y,
		Col: x,
	}
}

type Position struct {
	Row     int
	Col     int
	Nearest int
}

func (p1 Position) ManhattanDist(p2 Position) int {
	xOffset := float64(p1.Col - p2.Col)
	yOffset := float64(p1.Row - p2.Row)
	sum := math.Abs(xOffset) + math.Abs(yOffset)
	return int(sum)
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d) n=%d", p.Col, p.Row, p.Nearest)
}

func (p *Position) SetNearest(nearby Position) {
	p.Nearest = p.ManhattanDist(nearby)
}

// func chartSignal(p Position) [][]string {

// }

func updateMin(p1 Position, min Position) Position {
	if p1.Col < min.Col {
		min.Col = p1.Col
	}
	if p1.Row < min.Row {
		min.Row = p1.Row
	}
	return min
}

func updateMax(p1 Position, max Position) Position {
	if p1.Col > max.Col {
		max.Col = p1.Col
	}
	if p1.Row > max.Row {
		max.Row = p1.Row
	}
	return max
}

func updateMinMax(p1 Position, min Position, max Position) (Position, Position) {
	return updateMin(p1, min), updateMax(p1, max)
}

func blankGrid(min Position, max Position) [][]string {
	rows := max.Row - min.Row
	cols := max.Col - min.Col

	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = initRow(cols, ".")
	}
	return grid
}

func initRow(cols int, char string) []string {
	row := make([]string, cols)
	for i := 0; i < cols; i++ {
		row[i] = char
	}
	return row
}
