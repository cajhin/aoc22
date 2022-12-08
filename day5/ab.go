package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DEBUG bool = false
var Testmode bool = true
var fileLines []string

func main() {

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "prod" {
		Testmode = false
		fmt.Println("Production")
		ReadFileIntoGlobalArray("input")
	} else {
		fmt.Println("Test Run")
		ReadFileIntoGlobalArray("test")
	}

	//parse
	var moveN, moveF, moveT []int
	var stacks []string
	stacks_inited := false

	for _, line := range fileLines {
		if !stacks_inited {
			for i := 0; i < (len(line)+1)/4; i++ {
				stacks = append(stacks, "")
			}
			stacks_inited = true
		}
		if strings.ContainsRune(line, ']') {
			for i := 1; i < len(line); i++ {
				if line[i] == ']' {
					stacks[i/4] += string(line[i-1])
				}

			}
		} else if strings.Contains(line, "move") {
			s := strings.Split(line, " ")
			tmp1, _ := strconv.Atoi(s[1])
			tmp2, _ := strconv.Atoi(s[3])
			tmp3, _ := strconv.Atoi(s[5])
			moveN = append(moveN, tmp1)
			moveF = append(moveF, tmp2)
			moveT = append(moveT, tmp3)
		}

	}

	stacksA := SortStacksA(moveN, moveF, moveT, stacks)
	messageA := GetTopContainers(stacksA)
	stacksB := SortStacksB(moveN, moveF, moveT, stacks)
	messageB := GetTopContainers(stacksB)

	if Testmode {
		fmt.Println("Before:")
		for i, str := range stacks {
			fmt.Println("stacks", i, "=", str)
		}
		fmt.Println("Transformer:")
		fmt.Println(moveN)
		fmt.Println(moveF)
		fmt.Println(moveT)
	}

	fmt.Println("day5a solution: ", messageA, "is the top letter of each stack")
	fmt.Println("day5b solution: ", messageB, "is the top letter of each stack")
}

func GetTopContainers(stacks []string) string {
	var message string
	for _, s := range stacks {
		message += s[0:1]
	}

	return message
}

func SortStacksA(moveN []int, moveF []int, moveT []int, notouchStacks []string) []string {
	stacks := append(make([]string, 0, len(notouchStacks)), notouchStacks...)
	for i, _ := range moveN {
		num := moveN[i]
		from := moveF[i] - 1
		to := moveT[i] - 1
		for j := 0; j < num; j++ {
			stacks[to] = stacks[from][0:1] + stacks[to]
			stacks[from] = stacks[from][1:]
		}

		if DEBUG {
			fmt.Println("With", num, " ", from, " ", to)
			for i, str := range stacks {
				fmt.Println("stacks", i, "=", str)
			}
		}
	}

	return stacks
}

func SortStacksB(moveN []int, moveF []int, moveT []int, notouchStacks []string) []string {
	stacks := append(make([]string, 0, len(notouchStacks)), notouchStacks...)
	for i, _ := range moveN {
		num := moveN[i]
		from := moveF[i] - 1
		to := moveT[i] - 1
		stacks[to] = stacks[from][0:num] + stacks[to]
		stacks[from] = stacks[from][num:]

		if DEBUG {
			fmt.Println("With", num, " ", from, " ", to)
			for i, str := range stacks {
				fmt.Println("stacks", i, "=", str)
			}
		}
	}

	return stacks
}

func ReadFileIntoGlobalArray(filePath string) {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
}
