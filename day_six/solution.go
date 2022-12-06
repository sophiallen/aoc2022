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
	startPacketPos := startPacketMarkerPos(text)
	messageStartPos := startMessagePos(text)
	fmt.Printf("Start Packet: %+v\n", startPacketPos)
	fmt.Printf("Start message is at %+v", messageStartPos)

}

func startMessagePos(text string) int {
	i := 0
	textlen := len(text)
	for i = 0; i < len(text); i++ {
		nearEOL := i + 14 > textlen
		if nearEOL {
			fmt.Println("Hit eol searching for message")
			return -1
		}
		substr := text[i:i+14]
		isUnique := allUnique(substr)
		if isUnique {
			break
		}
	}
	return i + 14
}

func startPacketMarkerPos(text string) int {
	i := 0
	for i = 0; i < len(text); i++ {
		nearEOL := i + 4 > len(text)
		if (nearEOL) {
			fmt.Println("Hit EOL looking for start packet")
			break
		}
		subStr := text[i:i+4]
		isUnique := allUnique(subStr)
		if isUnique {
			break
		}
	}
	return i + 4
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