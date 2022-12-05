package day_two

import (
    "bufio"
    "fmt"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_two/data.txt"
}

const (
	scoreWin = 6
	scoreDraw = 3
	scoreLose = 0
	myRock = "X"
	myPaper = "Y"
	myScissors = "Z"
	stratLose = "X"
	stratDraw = "Y"
	stratWin = "Z"
	theirRock = "A"
	theirPaper = "B"
	theirScissors = "C"
)

var winningMoves = map[string]string {
	theirRock: myPaper,
	theirPaper: myScissors,
	theirScissors: myRock,
}
var drawMoves = map[string]string {
	theirRock: myRock,
	theirPaper: myPaper,
	theirScissors: myScissors,
}
var losingMoves = map[string]string {
	theirRock: myScissors,
	theirPaper: myRock,
	theirScissors: myPaper,
}
var shapeScores = map[string]int {
	myRock: 1,
	myPaper: 2,
	myScissors: 3,
}
var outcomeRock = map[string]int {
	myRock: scoreDraw,
	myPaper: scoreWin,
	myScissors: scoreLose,
}
var outcomePaper = map[string]int {
	myRock: scoreLose,
	myPaper: scoreDraw,
	myScissors: scoreWin,
}
var outcomeScissors = map[string]int {
	myPaper: scoreLose,
	myRock: scoreWin,
	myScissors: scoreDraw,
}
var outcomes = map[string]map[string]int {
	theirPaper: outcomePaper,
	theirRock: outcomeRock,
	theirScissors: outcomeScissors,
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	pt1Total := 0
	pt2Total := 0
	for scanner.Scan() {
		moves := getMoves(scanner) 
		if len(moves) == 0 {
			return
		}
		pt1Total += judgePt1(moves[0], moves[1])
		pt2Total += judgePt2(moves[0], moves[1])
	}
	fmt.Printf("Pt1 Score: %+v\nPt2 Score: %+v", pt1Total, pt2Total)
}

func getMoves(scanner *bufio.Scanner) []string {
	line := scanner.Text()
	moves := strings.Split(line, " ")
	if len(moves) == 0 {
		fmt.Println("ERR: NO MOVES!!")
		return moves
	}
	return moves
}
func judgePt1(theirmove string, myMove string) int { 
	return shapeScores[myMove] + outcomes[theirmove][myMove]
}

func judgePt2(theirMove string, myStrat string) int {
	score := stratToScore(myStrat)
	if myStrat == stratWin {
		return score + shapeScores[winningMoves[theirMove]]
	}
	if myStrat == stratLose {
		return score + shapeScores[losingMoves[theirMove]]
	}
	return score + shapeScores[drawMoves[theirMove]]
}

func stratToScore(strat string) int {
	if strat == stratWin {
		return scoreWin
	}
	if strat == stratDraw {
		return scoreDraw
	}
	return 0
}
