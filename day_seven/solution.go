package day_seven

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

type directory struct {
	Name     string
	Size     int
	Parent   string
	ChildIds []string
}

func (s Solution) GetDataPath() string {
	return "day_seven/data.txt"
}

func (s Solution) Solve(scanner *bufio.Scanner) {
	directories := map[string]directory{}

	directories["/"] = directory{
		Name: "/",
	}
	path := []string{}

	// for each line in the file
	for scanner.Scan() {
		// get line of text as a string
		line := scanner.Text()
		isCmd := isCommand(line)
		if isCmd {
			cmdParts := strings.Split(line, " ")
			if len(cmdParts) > 2 {
				if cmdParts[2] == ".." {
					_, path = pop(path)
				} else {
					path = append(path, cmdParts[2])
				}
			}
			continue
		}
		dirs := processLsLine(path, line, directories)
		directories = dirs

	}
	dirsToDelete := []int{}
	topDirSize := directories[path[0]].Size
	const spaceNeeded = 30000000
	const systemSpace = 70000000
	smallestRightSizedDir := systemSpace
	currentUnused := systemSpace - topDirSize
	for _, v := range directories {
		// fmt.Printf("Directory %+v size is %+v\n", v.Name, v.Size)
		if v.Size <= 100000 {
			dirsToDelete = append(dirsToDelete, v.Size)
		}
		if currentUnused+v.Size >= spaceNeeded {
			if v.Size < smallestRightSizedDir {
				smallestRightSizedDir = v.Size
			}
		}
	}
	reclaimableSpace := 0
	for _, size := range dirsToDelete {
		reclaimableSpace += size
	}

	fmt.Printf("Space to reclaim is %+v\n", reclaimableSpace)
	fmt.Printf("Smallest right-sized dir is %+v\n", smallestRightSizedDir)
}

func isCommand(text string) bool {
	return len(text) > 0 && text[0] == '$'
}

func pop(arr []string) (string, []string) {
	newLen := len(arr) - 1
	value := arr[newLen]
	newArr := arr[:newLen]
	return value, newArr
}

func processLsLine(path []string, text string, dirMap map[string]directory) map[string]directory {
	parts := strings.Split(text, " ")
	childPrefix := parts[0]
	parentName := strings.Join(path, ".")
	childName := parentName + "." + parts[1]
	parent := dirMap[parentName]

	if includesChild(parent.ChildIds, childName) {
		// child was already accounted for. Move on
		return dirMap
	}

	parent.ChildIds = append(parent.ChildIds, childName)

	if childPrefix == "dir" {
		dirMap = addDirectory(childName, parentName, dirMap)
	} else {
		// item is a file
		dirMap = addFile(childName, childPrefix, parentName, dirMap)
	}
	return dirMap
}

func addDirectory(dirName string, parentName string, dirMap map[string]directory) map[string]directory {

	parent := dirMap[parentName]
	dirMap[dirName] = directory{
		Parent: parent.Name,
		Name:   dirName,
	}
	parent.ChildIds = append(parent.ChildIds, dirName)
	dirMap[parent.Name] = parent

	return dirMap
}

func addFile(fileName string, sizeString string, parentName string, dirMap map[string]directory) map[string]directory {
	fileSize, err := strconv.Atoi(sizeString)
	if err != nil {
		fmt.Printf("ERROR!! with file %+v:\n%+v\n", fileName, err)
	}
	parent := dirMap[parentName]
	for len(parentName) > 0 {
		dirMap[parent.Name] = incrementSize(parent, fileSize)
		parent = dirMap[parent.Parent]
		parentName = parent.Name
	}
	return dirMap
}

func incrementSize(dir directory, size int) directory {
	dir.Size += size
	return dir
}

func includesChild(childIds []string, childName string) bool {
	for _, name := range childIds {
		if name == childName {
			return true
		}
	}
	return false
}
