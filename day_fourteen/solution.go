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

func (p Path) xtreme() int {
	rightmost := 0
	for _, p := range p.points {
		if p.X > rightmost {
			rightmost = p.X
		}
	}
	return rightmost
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
	return "day_fourteen/test.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	var farthestRight, farthestDown int
	paths := []Path{}

	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		fmt.Println(line)
		p := newPath(line)
		xt, yt := p.xtreme(), p.ytreme()
		if xt > farthestRight {
			farthestRight = p.xtreme()
		}
		if yt > farthestDown {
			farthestDown = yt
		}
		paths = append(paths, p)
	}
	fmt.Printf("farthest right %+v, farthest down %+v\n", farthestRight, farthestDown)

	grid := make([][]string, farthestDown)
	for i := 0; i > len(grid); i++ {
		grid[i] = make([]string, farthestRight)
	}

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
