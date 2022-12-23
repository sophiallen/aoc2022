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
	return "day_fifteen/data.txt"
}

const PT2_MIN = 0
const PT2_MAX = 4000000

type RowRange struct {
	min int
	max int
}

func (r RowRange) eclipses(r2 RowRange) bool {
	return r.min <= r2.min && r.max >= r2.max
}

func (r RowRange) contains(num int) bool {
	return r.min <= num && r.max >= num
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	// sensorPositions := []Position{}
	devices := map[string]bool{}
	rowRanges := make([][]RowRange, PT2_MAX+1)
	for i := range rowRanges {
		rowRanges[i] = []RowRange{}
	}
	// min := PT2_MAX / 2
	// max := PT2_MAX / 2
	// targetRow := 10
	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		positions := strings.Split(line, ":")
		sensorPos := lineToPosition(positions[0], "Sensor")
		beaconPos := lineToPosition(positions[1], "Beacon")

		devices[sensorPos.String()] = true
		devices[beaconPos.String()] = true

		sensorPos.SetNearest(beaconPos)
		// sensorPositions = append(sensorPositions, sensorPos)
		minRow := sensorPos.Row - sensorPos.Range
		maxRow := sensorPos.Row + sensorPos.Range
		if minRow < PT2_MIN {
			minRow = PT2_MIN
		}
		if maxRow > PT2_MAX {
			maxRow = PT2_MAX
		}
		for i := minRow; i <= maxRow; i++ {
			mi, mx := rangeInRow(sensorPos, i)
			rowRanges[i] = smush(rowRanges[i], RowRange{
				min: mi,
				max: mx,
			})
		}

		// if touchesRow(sensorPos, targetRow) {
		// 	mn, mx := rangeInRow(sensorPos, targetRow)
		// 	if mn < min {
		// 		min = mn
		// 	}
		// 	if mx > max {
		// 		max = mx
		// 	}
		// }
	}
	// possibles := 0
	// possibles = countPossible(min, max, targetRow, devices, sensorPositions)

	distressSignal := Position{}
	containGaps := 0
	for i, r := range rowRanges {
		if len(r) > 1 {
			containGaps++
			fmt.Printf("Num ranges is %+d\n", len(r))
			gap := findGap(r)
			if gap.min == gap.max {
				distressSignal.Col = gap.min
				distressSignal.Row = i
			}
		}
	}
	// fmt.Printf("Possibles: %d, min %d, max %d\n", possibles, min, max)
	fmt.Printf("Num gaps: %d\n", containGaps)
	fmt.Println("Distress signal:", distressSignal)
	fmt.Println("Tuning:", (distressSignal.Col*4000000)+distressSignal.Row)
}

func countPossible(min int, max int, targetRow int, devices map[string]bool, sensorPositions []Position) int {
	possibles := 0
	for i := min; i <= max; i++ {
		pos := Position{
			Row: targetRow,
			Col: i,
		}
		if devices[pos.String()] {
			continue
		}
		for s := 0; s < len(sensorPositions); s++ {
			if isInRange(pos, sensorPositions[s]) {
				possibles++
				break
			}
		}
	}
	return possibles
}

func lineToPosition(line string, thingType string) Position {
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
		Row:  y,
		Col:  x,
		Type: thingType,
	}
}

type Position struct {
	Row   int
	Col   int
	Range int
	Type  string
}

func (p1 Position) ManhattanDist(p2 Position) int {
	xOffset := float64(p1.Col - p2.Col)
	yOffset := float64(p1.Row - p2.Row)
	sum := math.Abs(xOffset) + math.Abs(yOffset)
	return int(sum)
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d) n=%d", p.Col, p.Row, p.Range)
}

func (p *Position) SetNearest(nearby Position) {
	p.Range = p.ManhattanDist(nearby)
}

func isInRange(pointA Position, sensorPos Position) bool {
	return sensorPos.ManhattanDist(pointA) <= sensorPos.Range
}

func touchesRow(sensor Position, row int) bool {
	return sensor.Row+sensor.Range >= row && sensor.Row-sensor.Range <= row
}

func rangeInRow(sensor Position, row int) (int, int) {
	intOffset := sensor.Row - row
	offSet := math.Abs(float64(intOffset))
	min := (sensor.Col - sensor.Range) + int(offSet)
	max := (sensor.Col + sensor.Range) - int(offSet)
	return min, max
}

func combineRanges(r1 RowRange, r2 RowRange) []RowRange {
	if r1.eclipses(r2) {
		return []RowRange{r1}
	}
	if r2.eclipses(r1) {
		return []RowRange{r2}
	}
	// overlap with r1 on low end
	if r1.max >= r2.min && r1.min < r2.min {
		return []RowRange{
			{
				min: r1.min,
				max: r2.max,
			},
		}
	}
	// overlap with r2 on low end
	if r2.max >= r1.min && r2.min < r1.min {
		return []RowRange{
			{
				min: r2.min,
				max: r1.max,
			},
		}
	}
	// touching borders
	if r1.max+1 == r2.min {
		return []RowRange{
			{
				min: r1.min,
				max: r2.max,
			},
		}
	}
	if r2.max+1 == r1.min {
		return []RowRange{
			{
				min: r2.min,
				max: r1.max,
			},
		}
	}

	// there's a gap
	return []RowRange{r1, r2}
}

func reduceRanges(ranges []RowRange) []RowRange {
	i := 0
	for i+1 < len(ranges) {
		r1 := ranges[i]
		r2 := ranges[i+1]
		combined := combineRanges(r1, r2)
		if len(combined) == 1 {
			newRanges := append(ranges[:i], combined[0])
			newRanges = append(newRanges, ranges[i+2:]...)
			ranges = newRanges
			continue
		}
		i++
	}
	return ranges
}

func smush(current []RowRange, insert RowRange) []RowRange {
	newRanges := make([]RowRange, len(current))
	didCombine := false
	for i, r := range current {
		combined := combineRanges(insert, r)
		if len(combined) == 1 {
			newRanges[i] = combined[0]
			didCombine = true
			continue
		}
		newRanges[i] = r
	}
	if didCombine {
		reduced := reduceRanges(newRanges)
		for len(reduced) < len(newRanges) {
			newRanges = reduced
			reduced = reduceRanges(newRanges)
		}
		return reduced
	}
	return append(current, insert)
}

func findGap(ranges []RowRange) RowRange {
	r1 := ranges[0]
	r2 := ranges[1]
	if r1.max < r2.min {
		return RowRange{
			min: r1.max + 1,
			max: r2.min - 1,
		}
	}
	if r2.max < r1.min {
		return RowRange{
			min: r2.max + 1,
			max: r2.min - 1,
		}
	}
	return RowRange{}
}
