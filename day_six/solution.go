package day_six

import (
	"bufio"
	"fmt"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_six/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {

	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	startPacketPos := getPostionLastUnique(4, text)
	messageStartPos := getPostionLastUnique(14, text)
	fmt.Printf("Start Packet: %+v\n", startPacketPos)
	fmt.Printf("Start message is at %+v", messageStartPos)

}

func getPostionLastUnique(lenUnique int, text string) int {
	i := 0
	textlen := len(text)
	for i = 0; i < len(text); i++ {
		nearEOL := i+lenUnique > textlen
		if nearEOL {
			fmt.Println("Hit eol searching for message")
			return -1
		}
		substr := text[i : i+lenUnique]
		isUnique := allUnique(substr)
		if isUnique {
			break
		}
	}
	return i + lenUnique
}

func allUnique(substr string) bool {
	uniques := map[rune]bool{}
	for _, r := range substr {
		if uniques[r] {
			return false
		}
		uniques[r] = true
	}
	return true
}
