package day_three

import (
	"fmt"
	"bufio"

	// "strings"
)

type Solution struct {}

func (s Solution) GetDataPath() string {
	return "day_three/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	var (
		sumOfTypes = 0
		countGroup = 0
		badgeSum = 0
		curGroup = make([]string, 3)
	)
	// sumOfTypes := 0
	// curGroup := make([]string, 3)
	// countGroup := 0
	// badgeSum := 0
	for scanner.Scan() {
		text := scanner.Text()
		// pt 1
		sumOfTypes += getRucksackPriority(text)
	
		// part two
		idx := countGroup % 3
		curGroup[idx] = text
		if idx == 2 {
			badge := getBadge(curGroup)
			badgeSum += getPriority(badge)
		}
		countGroup += 1
	}
	fmt.Printf("Sum of Types is %+v", sumOfTypes)
}

func getRucksackPriority(text string) int {
	first, second := splitLine(text)
	common := findCommonItem(first, second)
	priority := getPriority(common)
	return priority
}

// func solvePartTwo(scanner *bufio.Scanner) {
// 	curGroup := make([]string, 3)
// 	countGroup := 0
// 	badgeSum := 0
// 	for scanner.Scan() {
// 		idx := countGroup % 3
// 		curGroup[idx] = scanner.Text()
// 		if idx == 2 {
// 			badge := getBadge(curGroup)
// 			fmt.Printf("Badge is %+v\n", string(badge))
// 			badgeSum += getPriority(badge)
// 		}
// 		countGroup += 1
// 	}
// 	fmt.Printf("\nSum is %+v\n", badgeSum)
// }

func getBadge(rucksacks []string) rune {
	accum := map[rune]int{}
	for _, sack := range(rucksacks) {
		accum = incrementUnique(sack, accum)
	}
	for letter, count := range accum {
		if count == 3 {
			return letter
		}
	}
	return '?'
}

func incrementUnique(list string, accum map[rune]int) map[rune]int {
	uniques := map[rune]int{}
	for _, letter := range list {
		if uniques[letter] > 0 {
			continue
		}
		uniques[letter] = 1
		accum[letter] = accum[letter] + 1
	}
	return accum
}

func splitLine(text string) (string, string) {
	halfLen := len(text) / 2
	firsthalf := text[:halfLen]
	sechalf := text[halfLen:]
	return firsthalf, sechalf
}

func findCommonItem(first string, second string) rune {
	runeMap := map[rune]bool{}
	for _, letter := range(first) {
		runeMap[letter] = true
	}
	for _, letter := range(second) {
		if runeMap[letter] {
			return letter
		}
	}
	return '?'
}

func getPriority(letter rune) int {
	val := int(letter)
	if (letter >= 97) {
		return val - 96
	}
	return val - 38
}