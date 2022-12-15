package day_thirteen

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_thirteen/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	// for each line in the file
	pairs := [][]PacketPiece{}
	curPair := []PacketPiece{}
	div1, _ := parsePacket("[[2]]", 1)
	div2, _ := parsePacket("[[6]]", 1)
	flatList := byElfOrder{div1, div2}

	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		// fmt.Println(line)
		if len(line) == 0 {
			pairs = append(pairs, curPair)
			curPair = []PacketPiece{}
			continue
		}
		nextPiece, completed := parsePacket(line, 1)
		if completed != len(line)-1 {
			fmt.Printf("Completed piece early?? %+v %+v %+v\n", completed, nextPiece, len(line))
		}
		flatList = append(flatList, nextPiece)
		// fmt.Printf("Matches = %+v\n", matches)
		curPair = append(curPair, nextPiece)
	}
	pairs = append(pairs, curPair)
	correctPairTotal := 0
	for i, pair := range pairs {
		comparison := compareLists(pair[0], pair[1])
		// fmt.Printf("Pair %+v is %+v\n", i+1, comparison)
		if comparison >= 0 {
			correctPairTotal += i + 1
		}
	}
	// last := pairs[len(pairs)-1]
	// fmt.Printf("Last pair was %+v and %+v\n", last[0].slice, last[1].slice)
	fmt.Printf("Sum of correct indices:%+v\n", correctPairTotal)

	// fmt.Println("----Unsorted-----")
	// for _, v := range flatList {
	// 	fmt.Println(v.String())
	// }
	// fmt.Println("\n----Sorted-----")
	sort.Sort(flatList)
	indices := []int{}
	for i, v := range flatList {
		stringRep := v.String()
		if stringRep == "[[2]]" || stringRep == "[[6]]" {
			indices = append(indices, i+1)
		}
	}

	fmt.Printf("Indices are %+v\n", indices)
	fmt.Printf("Multiplied is %+v\n", indices[0]*indices[1])
}

func createDivider(num int) PacketPiece {
	inner := PacketPiece{
		num: &num,
	}
	middle := PacketPiece{
		slice: &[]PacketPiece{inner},
	}
	outer := PacketPiece{
		slice: &[]PacketPiece{middle},
	}
	return outer
}

func strToNum(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Numstr did not convert! \"%+s\", %d\n", numStr, num)
	}
	return num
}

func compareLists(leftPiece, rightPiece PacketPiece) int {
	left := leftPiece.getList()
	right := rightPiece.getList()

	// fmt.Printf("Compare slices %+v AND %+v\n", *left, *right)

	for i, l := range left {
		if i == len(right) {
			// fmt.Printf("Right is shorter than left, out of order\n")
			// right ran out of items before left
			return -1
		}
		rv := right
		r := rv[i]

		if l.isInt() && r.isInt() {
			// fmt.Printf("Compare %d vs %d\n", *l.num, *r.num)
			diff := r.getInt() - l.getInt()
			if diff > 0 {
				return 1
			}
			if diff == 0 {
				continue
			}
			return -1
		}
		if l.isArr() && r.isArr() {
			correct := compareLists(l, r)
			if correct == 0 {
				continue
			}
			return correct
			// if equal, continue
			// if definitively correct, return correct.
		}
		// mixed set
		if l.isInt() {
			// fmt.Printf("Wrapping %+v in list\n", l.getInt())
			// wrap l in list
			ln := PacketPiece{
				slice: &[]PacketPiece{l},
			}
			correct := compareLists(ln, r)
			if correct == 0 {
				continue
			}
			return correct
		}
		// r must be int, wrap in slice
		rn := PacketPiece{
			slice: &[]PacketPiece{r},
		}
		correct := compareLists(l, rn)
		if correct == 0 {
			continue
		}
		return correct

	}
	if len(right) > len(left) {
		// fmt.Printf("left is shorter\n")
		return 1
	}
	// fmt.Println("Lists are equal")
	return 0
}
