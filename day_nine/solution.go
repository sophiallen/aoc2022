package day_nine

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

type position struct {
	X int
	Y int
}

func (p position) toString() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (s Solution) GetDataPath() string {
	return "day_nine/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	// for each line in the file
	headPos := position{}
	tailPos := position{}

	longerRope := make([]position, 9)

	// coordsVisited := []string{tailPos.toString()}
	uniqueVisitsP1 := map[string]bool{}
	uniqueVisitsP2 := map[string]bool{}
	uniqueVisitsP1[tailPos.toString()] = true

	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		parts := strings.Split(line, " ")
		times, _ := strconv.Atoi(parts[1])
		if parts[0] == "R" {
			for i := 0; i < times; i++ {
				headPos.X = headPos.X + 1
				longerRope = pullRope(headPos, longerRope)
				uniqueVisitsP1[longerRope[0].toString()] = true
				uniqueVisitsP2[longerRope[8].toString()] = true
			}
		}
		if parts[0] == "U" {
			for i := 0; i < times; i++ {
				headPos.Y = headPos.Y + 1
				longerRope = pullRope(headPos, longerRope)
				uniqueVisitsP1[longerRope[0].toString()] = true
				uniqueVisitsP2[longerRope[8].toString()] = true
			}
		}
		if parts[0] == "D" {
			for i := 0; i < times; i++ {
				headPos.Y = headPos.Y - 1
				longerRope = pullRope(headPos, longerRope)
				uniqueVisitsP1[longerRope[0].toString()] = true
				uniqueVisitsP2[longerRope[8].toString()] = true
			}
		}
		if parts[0] == "L" {
			for i := 0; i < times; i++ {
				headPos.X = headPos.X - 1
				longerRope = pullRope(headPos, longerRope)
				uniqueVisitsP1[longerRope[0].toString()] = true
				uniqueVisitsP2[longerRope[8].toString()] = true
			}
		}
	}
	fmt.Printf("Unique coordinates %+v, %+v", len(uniqueVisitsP1), len(uniqueVisitsP2))

}

func pullRope(headPos position, following []position) []position {
	newPositions := make([]position, 9)
	head := headPos
	for i, knot := range following {
		newPositions[i] = follow(head, knot)
		head = newPositions[i]
	}
	return newPositions
}

func follow(hPos position, tPos position) position {
	offsetX := hPos.X - tPos.X
	offsetY := hPos.Y - tPos.Y

	if offsetY > 1 {
		tPos.Y = tPos.Y + 1
		if offsetX > 0 {
			tPos.X = tPos.X + 1
		}
		if offsetX < 0 {
			tPos.X = tPos.X - 1
		}
		return tPos
	}

	if offsetY < -1 {
		tPos.Y = tPos.Y - 1
		if offsetX > 0 {
			tPos.X = tPos.X + 1
		}
		if offsetX < 0 {
			tPos.X = tPos.X - 1
		}
		return tPos
	}

	if offsetX > 1 {
		tPos.X = tPos.X + 1
		if offsetY > 0 {
			tPos.Y = tPos.Y + 1
		}
		if offsetY < 0 {
			tPos.Y = tPos.Y - 1
		}
		return tPos
	}

	if offsetX < -1 {
		tPos.X = tPos.X - 1
		if offsetY > 0 {
			tPos.Y = tPos.Y + 1
		}
		if offsetY < 0 {
			tPos.Y = tPos.Y - 1
		}
	}

	return tPos
}
