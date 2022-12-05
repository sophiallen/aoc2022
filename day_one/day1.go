package day_one

import (
    "bufio"
    "fmt"
	"strconv"
)

func (s Solution) GetDataPath() string {
	return "day_one/data.txt"
}

type Solution struct {}

func (s Solution) Solve(scanner *bufio.Scanner) {
  	mostCals := 0
	curElfTotalCals := 0
    leaderBoard := []int{0,0,0}

    for scanner.Scan() {
		calString := scanner.Text()
		if len(calString) == 0 {
			if curElfTotalCals > mostCals {
				mostCals = curElfTotalCals
			}
			leaderBoard = checkLeaderBoard(curElfTotalCals, leaderBoard)
			curElfTotalCals = 0
			continue
		}
		calInt, err := strconv.Atoi(calString)
		if (err != nil) {
			fmt.Printf("Could not parse %+v, err %+v\n", calString, err)
		}
		curElfTotalCals += calInt
	}
	checkLeaderBoard(curElfTotalCals, leaderBoard)
	leadsTotal := 0
	for _, cur := range(leaderBoard) {
		leadsTotal += cur
	}
	fmt.Printf("mostCals is %+v\nleads total cals: %+v\n", mostCals, leadsTotal)
}

func checkLeaderBoard(newScore int, leaders []int) []int {
	current := newScore 
	prevLeader := 0
	for i := 0; i < len(leaders); i++ {
		if current > leaders[i] {
			prevLeader = leaders[i]
			leaders[i] = current
			current = prevLeader
		}
	}
	return leaders
}
