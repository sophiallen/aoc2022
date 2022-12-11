package day_eleven

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) GetDataPath() string {
	return "day_eleven/data.txt"
}

var globalMod int64 = 1

type Monkey struct {
	Items        []int64
	ID           int
	CatchCount   int
	InspectCount int
	TestMod      int64
	targets      []int
	op           operation
}

type operation struct {
	verb    string
	subject int64
}

func (m *Monkey) inspect(item int64) int64 {
	m.InspectCount += 1
	var newVal int64
	subject := m.op.subject
	if subject < 0 {
		subject = item
	}
	if strings.Contains(m.op.verb, "*") {
		newVal = item * subject
	} else {
		newVal = item + subject
	}

	return newVal % globalMod
}

func (m *Monkey) throwItem(item int64) int {
	m.Items = m.Items[1:]
	mod := item % m.TestMod
	if mod == 0 {
		return m.targets[0]
	} else {
		return m.targets[1]
	}
}

func (m *Monkey) catchItem(item int64) {
	m.Items = append(m.Items, item)
	m.CatchCount += 1
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	curMonkey := 0
	monkeys := []Monkey{}
	items := [][]int64{}
	mods := []int64{}
	targets := [][]int{}
	ops := []operation{}

	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		if len(line) == 0 {
			curMonkey++
		}
		if strings.Contains(line, "Monkey") {
			monkeys = append(monkeys, Monkey{ID: curMonkey})
		}
		if strings.Contains(line, "Starting") {
			list := []int64{}
			h2 := strings.Split(line, ": ")
			nums := strings.Split(h2[1], ", ")
			for _, i := range nums {
				num, err := strconv.Atoi(i)
				num64 := int64(num)
				if err != nil {
					fmt.Println("ERROR READING START LIST")
					fmt.Println(num)
				}
				list = append(list, num64)
			}
			items = append(items, list)
			continue
		}
		if strings.Contains(line, "Operation") {
			o := operation{}
			fmt.Printf("Checking operation %+v\n", line)
			if strings.Contains(line, "*") {
				o.verb = "*"
			} else {
				o.verb = "+"
			}
			n := getLastAsInt64(line)
			o.subject = n
			fmt.Printf("Resulting op for %+v is %+v %+v\n", curMonkey, o.verb, o.subject)
			ops = append(ops, o)
			continue
		}
		if strings.Contains(line, "Test") {
			m := getLastAsInt64(line)
			mods = append(mods, m)
			globalMod *= m
			continue
		}
		if strings.Contains(line, "true") {
			t := []int{}
			tg := getLastAsInt(line)
			t = append(t, tg)
			targets = append(targets, t)
			continue
		}
		if strings.Contains(line, "false") {
			t := targets[curMonkey]
			tg := getLastAsInt(line)
			t = append(t, tg)
			targets[curMonkey] = t
		}
	}

	for id, m := range monkeys {
		m.ID = id
		m.Items = items[id]
		m.TestMod = mods[id]
		m.targets = targets[id]
		m.op = ops[id]
		monkeys[id] = m
	}

	for round := 0; round < 10000; round++ {
		for i, monkey := range monkeys {
			// fmt.Printf("Monkey %+v's start items: %+v\n", monkey.ID, monkey.Items)
			for _, item := range monkey.Items {
				itemVal := monkey.inspect(item)
				target := monkey.throwItem(itemVal)
				monkeys[target].catchItem(itemVal)
			}
			monkeys[i] = monkey
		}
	}

	highest := 0
	secHighest := 0

	for _, m := range monkeys {
		fmt.Printf("Monkey %+v inspect count is %+v\n", m.ID, m.InspectCount)
		if m.CatchCount > highest {
			secHighest = highest
			highest = m.InspectCount
			continue
		}
		if m.CatchCount > secHighest {
			secHighest = m.InspectCount
		}
	}

	fmt.Printf("Highest is %+v, SecHighest is %+v, monkeyBiz is %+v", highest, secHighest, highest*secHighest)

}

func getLastAsInt64(str string) int64 {
	s := strings.Split(str, " ")
	n, e := strconv.Atoi(s[len(s)-1])
	if e != nil {
		return -1
	}
	return int64(n)
}

func getLastAsInt(str string) int {
	s := strings.Split(str, " ")
	n, e := strconv.Atoi(s[len(s)-1])
	if e != nil {
		return -1
	}
	return n
}

// func monkeyWithItems(items string) Monkey {
// 	startItems := strings.Split(items, ", ")
// 	itemArr := make([]int, len(startItems))
// 	for i, item := range startItems {
// 		itemArr[i], _ = strconv.Atoi(item)
// 	}
// 	return Monkey{
// 		items: itemArr,
// 	}
// }

// func setMonkeyOperation(operation string, monkey Monkey) Monkey {

// }
