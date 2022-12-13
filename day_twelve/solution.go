package day_twelve

import (
	"bufio"
	"fmt"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_twelve/data.txt"
}

type position struct {
	col int
	row int
}

func (p position) toString() string {
	return fmt.Sprintf("(%+v, %+v)", p.col, p.row)
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	heightMap := [][]rune{}
	startPos := position{}
	dest := position{}
	curRow := 0
	alisters := []position{}
	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		arr := []rune(line)
		s := strings.IndexRune(line, 'S')
		if s >= 0 {
			startPos.col = s
			startPos.row = curRow
			arr[s] = 'a'
			alisters = append(alisters, startPos)
		}
		d := strings.IndexRune(line, 'E')
		if d >= 0 {
			dest.col = d
			dest.row = curRow
			arr[d] = 'z'
		}
		for i, l := range arr {
			if l == 'a' {
				alisters = append(alisters, position{
					row: curRow,
					col: i,
				})
			}
		}
		heightMap = append(heightMap, arr)
		curRow++
	}
	fmt.Printf("There are %+v rows with len %+v\n", len(heightMap), len(heightMap[0]))
	fmt.Printf("Start pos is %+v\nDest is %+v\n", startPos.toString(), dest.toString())

	pt1Steps, _ := howManySteps(startPos, dest, heightMap)
	leastSteps := pt1Steps

	for _, pos := range alisters {
		steps, found := howManySteps(pos, dest, heightMap)
		if found && steps < leastSteps {
			leastSteps = steps
		}
	}
	fmt.Printf("Steps taken: %+v\n", pt1Steps)
	fmt.Printf("Least steps possible: %d", leastSteps)
}

func stepForward(paths []position, hm [][]rune, destination position, visited map[string]bool) ([]position, bool, map[string]bool) {
	nextStep := []position{}
	for _, path := range paths {
		if visited[path.toString()] {
			continue
		}
		trails := findPathways(path, hm)

		for _, dest := range trails {
			if dest.col == destination.col && dest.row == destination.row {
				return []position{}, true, visited
			}
			nextStep = append(nextStep, dest)
		}
		visited[path.toString()] = true
	}
	return nextStep, false, visited
}

func howManySteps(startPos position, goal position, hmap [][]rune) (int, bool) {
	paths := findPathways(startPos, hmap)
	arrived := false
	steps := 0
	visited := map[string]bool{}
	visited[startPos.toString()] = true

	for !arrived && len(paths) > 0 {
		steps += 1

		paths, arrived, visited = stepForward(paths, hmap, goal, visited)
		if arrived {
			steps += 1
			break
		}
	}
	return steps, arrived
}

func findPathways(startPos position, h [][]rune) []position {
	exits := []position{}
	cur := h[startPos.row][startPos.col]
	if cur == 'S' {
		cur = 'a'
	}
	// up
	if startPos.row > 0 {
		above := h[startPos.row-1][startPos.col]
		if isWalkable(cur, above) {
			exits = append(exits, position{col: startPos.col, row: startPos.row - 1})
		}
	}
	// left
	if startPos.col > 0 {
		left := h[startPos.row][startPos.col-1]
		if isWalkable(cur, left) {
			exits = append(exits, position{col: startPos.col - 1, row: startPos.row})
		}
	}
	// right
	if startPos.col+1 < len(h[startPos.row]) {
		right := h[startPos.row][startPos.col+1]
		if isWalkable(cur, right) {
			exits = append(exits, position{row: startPos.row, col: startPos.col + 1})
		}
	}
	// down
	if startPos.row+1 < len(h) {
		down := h[startPos.row+1][startPos.col]
		if isWalkable(cur, down) {
			exits = append(exits, position{col: startPos.col, row: startPos.row + 1})
		}
	}
	return exits
}

func isWalkable(a rune, b rune) bool {
	if int(a) > int(b) {
		return true
	}
	return int(b)-int(a) <= 1
}
