package day_four

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Solution struct {}

func (s Solution) GetDataPath() string {
	return "day_four/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	sumPairsContained := 0
	pairsOverlap := 0

	for scanner.Scan() {
		line := scanner.Text()
		pairs := splitPairs(line)
		set1 := getBounds(pairs[0])
		set2 := getBounds(pairs[1])
		if (completelyContains(set1, set2) || completelyContains(set2, set1)) {
			sumPairsContained += 1
			pairsOverlap += 1
			continue
		}
		if overLaps(set1, set2) {
			fmt.Printf("%+v overlaps %+v\n", set1, set2)
			pairsOverlap += 1
		}
	}
	fmt.Printf("Pairs contained: %+v\n Pairs overlap: %+v", sumPairsContained, pairsOverlap)
}

func splitPairs(line string) []string {
	return strings.Split(line, ",")
}

func getBounds(numstr string) []int {
	bounds := strings.Split(numstr,"-")
	res := make([]int, 2)
	for i,n := range bounds {
		num, _ := strconv.Atoi(n)
		res[i] = num
	}
	return res
}

func completelyContains(r1 []int, r2 []int) bool {
	return (r1[0] >= r2[0]  && r1[1] <= r2[1])
}
func overLaps(r1 []int, r2 []int) bool {
	r1Lower := r1[0]
	r1Upper := r1[1]
	r2Lower := r2[0]
	r2Upper := r2[1] 
	if r1Lower <= r2Lower {
		return r1Upper >= r2Lower 
	}
	// else r1 lower is greater than r2 lower
	if r1Lower >= r2Lower {
		return r1Lower <= r2Upper
	}
	return false
}