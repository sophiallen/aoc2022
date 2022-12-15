package day_thirteen

import (
	"fmt"
	"strings"
	"unicode"
)

type PacketPiece struct {
	num   *int
	slice *[]PacketPiece
}

func parsePacket(line string, start int) (PacketPiece, int) {
	piece := PacketPiece{}
	slice := []PacketPiece{}
	numStr := ""
	for i := start; i < len(line); i++ {
		if unicode.IsNumber(rune(line[i])) {
			numStr += string(line[i])
			continue
		}
		if line[i] == ',' && len(numStr) > 0 {
			num := strToNum(numStr)
			slice = append(slice, PacketPiece{
				num: &num,
			})
			numStr = ""
			continue
		}
		if line[i] == ']' {
			if len(numStr) > 0 {
				num := strToNum(numStr)
				slice = append(slice, PacketPiece{
					num: &num,
				})
			}

			piece.slice = &slice
			return piece, i
		}
		if line[i] == '[' {
			nextPiece, pieceEnd := parsePacket(line, i+1)
			slice = append(slice, nextPiece)
			i = pieceEnd
		}
	}
	piece.slice = &slice
	return piece, -1
}

func (p PacketPiece) isInt() bool {
	return p.num != nil
}
func (p PacketPiece) isArr() bool {
	return p.slice != nil
}

func (p PacketPiece) getInt() int {
	if p.num == nil {
		return -1
	}
	return *p.num
}

func (p PacketPiece) getList() []PacketPiece {
	if p.slice != nil {
		return *p.slice
	}
	return []PacketPiece{}
}

func (p PacketPiece) String() string {
	if p.num != nil {
		// n := *p.num
		return fmt.Sprintf("%d", *p.num)
	}
	if p.slice != nil {
		strs := []string{}
		for _, pp := range *p.slice {
			strs = append(strs, pp.String())
		}
		return "[" + strings.Join(strs, ",") + "]"
	}
	return "wtf"
}

type byElfOrder []PacketPiece

func (beo byElfOrder) Len() int {
	return len(beo)
}

func (beo byElfOrder) Swap(i, j int) {
	beo[i], beo[j] = beo[j], beo[i]
}

func (beo byElfOrder) Less(i, j int) bool {
	left := beo[i]
	right := beo[j]
	if left.isArr() && right.isArr() {
		return compareLists(left, right) >= 0
	}
	fmt.Printf("!! Tried to sort pieces that were not slices!! D: \n\n")
	return false
}
