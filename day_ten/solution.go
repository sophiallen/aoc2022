package day_ten

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_ten/data.txt"
}

type cycle struct {
	value int
	op    string
	pixel string
}

// PT 2
// Sprites are 3 px wide
// X controls horizontal position of middle of sprite
// if sprites position is where crt is currently drawing, draw the pixel

func (s Solution) Solve(scanner *bufio.Scanner) {
	cycles := []cycle{}
	curCycle := 0
	curRegister := 1
	// rows := []string{}
	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		c := cycle{
			value: curRegister,
			op:    line,
			pixel: toPixel(curRegister, xpos(curCycle)),
		}
		cycles = append(cycles, c)

		if line == "noop" {
			// end cycle, proceed to next.
			curCycle++
			continue
		}

		parts := strings.Split(line, " ")
		increment, _ := strconv.Atoi(parts[1])

		curCycle++
		c1 := cycle{
			op:    line,
			value: curCycle,
			pixel: toPixel(curRegister, xpos(curCycle)),
		}
		cycles = append(cycles, c1)
		curRegister += increment
		curCycle++
	}

	cycles = append(cycles, cycle{value: curRegister})
	runningStrength := 0
	row := ""
	for i := 0; i < len(cycles); i++ {
		fmt.Printf(cycles[i].pixel)
		if (i+1)%40 == 0 {
			fmt.Printf("\n")
		}
		runningStrength += (i + 1) * cycles[i].value
	}
	fmt.Println(row)
	// fmt.Printf("Final strength: %+v", runningStrength)
}

func xpos(cycle int) int {
	row := cycle / 40
	return cycle - (row * 40)
}

func toPixel(register int, xpos int) string {
	offset := float32(xpos - register)
	if math.Abs(float64(offset)) < 2 {
		return "#"
	}
	return "."
}
