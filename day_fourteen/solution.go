package day_fourteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

type Point struct {
	X int
	Y int
}

type Path struct {
	points []Point
}

const DROP_ZONE = 500

func (p Path) xtremes() (int, int) {
	rightmost := 0
	leftmost := p.points[0].X
	for _, p := range p.points {
		if p.X > rightmost {
			rightmost = p.X
		}
		if p.Y < leftmost {
			leftmost = p.X
		}
	}
	return rightmost, leftmost
}

func (p Path) ytreme() int {
	deepest := 0
	for _, p := range p.points {
		if p.Y > deepest {
			deepest = p.Y
		}
	}
	return deepest
}

func (s Solution) GetDataPath() string {
	return "day_fourteen/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	var farthestRight, farthestLeft, farthestDown int
	paths := []Path{}

	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		p := newPath(line)
		paths = append(paths, p)
	}

	_, farthestLeft = paths[0].xtremes()

	for _, p := range paths {
		maxX, minX := p.xtremes()
		yt := p.ytreme()
		if minX < farthestLeft {
			farthestLeft = minX
		}
		if maxX > farthestRight {
			farthestRight = maxX
		}
		if yt > farthestDown {
			farthestDown = yt
		}
	}
	fmt.Printf("farthest right %+v, farthest down %+v, farthest Left %+v\n", farthestRight, farthestDown, farthestLeft)

	grid := make([][]string, farthestDown+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = initRow(farthestRight + 1)
	}

	for _, p := range paths {
		grid = chartPath(p, grid)
	}

	voided := false
	grainsOfSand := 0
	for !voided {
		grid, voided = dropSand(grid, false)
		grainsOfSand++
	}
	dropSand(grid, true)

	trimmedGrid := make([][]string, farthestDown+1)
	for i := range trimmedGrid {
		trimmedGrid[i] = grid[i][farthestLeft:]
	}
	printGrid(trimmedGrid)

	fmt.Println("Fell into void:", voided)
	fmt.Println("Grains of Sand:", grainsOfSand-1)

}

func newPath(line string) Path {
	pts := strings.Split(line, " -> ")
	pointlist := make([]Point, len(pts))
	for i, pt := range pts {
		pointlist[i] = newPoint(pt)
	}
	return Path{
		points: pointlist,
	}
}

func newPoint(str string) Point {
	parts := strings.Split(str, ",")
	x, e := strconv.Atoi(parts[0])
	if e != nil {
		fmt.Printf("Invalid x value %+v\n", x)
		return Point{}
	}
	y, e := strconv.Atoi(parts[1])
	if e != nil {
		fmt.Printf("Invalid y value %+v\n", x)
		return Point{}
	}

	return Point{
		X: x,
		Y: y,
	}
}

func fillBetween(p1 Point, p2 Point, grid [][]string) [][]string {
	diffX := p2.X - p1.X
	dirX := 1
	diffY := p2.Y - p1.Y
	dirY := 1

	if diffX < 0 {
		dirX = -1
		diffX *= -1
	}
	if diffY < 0 {
		dirY = -1
		diffY *= -1
	}

	if diffX != 0 {
		row := grid[p1.Y]
		for i := 0; i < diffX; i++ {
			x := p1.X + (i * dirX)
			row[x] = "#"
		}
		grid[p1.Y] = row
	}

	if diffY != 0 {
		for i := 0; i < diffY; i++ {
			rn := p1.Y + (i * dirY)
			grid[rn][p1.X] = "#"
		}
	}
	return grid
}

func chartPath(p Path, grid [][]string) [][]string {
	pathlen := len(p.points)
	for i := 0; i < pathlen; i++ {
		p1 := p.points[i]
		grid[p1.Y][p1.X] = "#"
		if i+1 < pathlen {
			p2 := p.points[i+1]
			grid = fillBetween(p1, p2, grid)
		}
	}

	return grid
}

func initRow(size int) []string {
	row := make([]string, size)
	for i := range row {
		row[i] = "."
	}
	return row
}

func dropSand(grid [][]string, trace bool) ([][]string, bool) {
	fellIntoVoid := false
	x, y := DROP_ZONE, 0

	for y < len(grid) {
		// fmt.Printf("Sand is at (%+v, %+v)\n", x, y)
		if !isOccupied(grid[y][x]) {
			// fmt.Println("Moving down")
			if trace {
				grid[y][x] = "~"
			}
			y++
			continue
		}
		if x-1 >= 0 && !isOccupied(grid[y][x-1]) {
			// fmt.Println("Diagonal left")
			x = x - 1
			if trace {
				grid[y][x] = "~"
			}
			y++
			continue
		}
		if x+1 < len(grid[y]) && !isOccupied(grid[y][x+1]) {
			// fmt.Println("Diagonal right")
			x = x + 1
			if trace {
				grid[y][x] = "~"
			}
			y++
			continue
		}

		// fmt.Println("Hit bottom, backing up one and stopping")
		y--
		break
	}
	fellIntoVoid = x >= len(grid[0]) || y >= len(grid)
	if !fellIntoVoid {
		grid[y][x] = "o"
	}
	return grid, fellIntoVoid
}

func isOccupied(space string) bool {
	return space == "#" || space == "o"
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		r := strings.Join(row, "")
		fmt.Println(r)
	}
}
