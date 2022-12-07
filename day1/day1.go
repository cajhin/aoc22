package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//io
	filePath := "input"
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	//parse
	calories := []int{0}
	elfidx := 0
	for _, line := range fileLines {
		if line == "" {
			elfidx++
			calories = append(calories, 0)
		} else {
			intVar, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("bad line: ", line)
			} else {
				calories[elfidx] += intVar
			}
		}
	}

	//Solution A:
	maxidxA, maxvalA := getTopScore(calories)
	fmt.Println("day1a solution: the top elf #", maxidxA+1, " carries", maxvalA, "calories")

	//Solution B: get top three
	topThreeAdded := 0

	for i := 0; i < 3; i++ {
		maxidx, maxval := getTopScore(calories)
		fmt.Println("Medal", i+1, " goes to elf #", maxidx+1, "of", len(calories)+1, "with", maxval, "calories")
		topThreeAdded += maxval
		calories[maxidx] = 0
	}

	fmt.Println("day1b solution: the top three elves together carry", topThreeAdded)
}

func getTopScore(arr []int) (int, int) {
	maxVal := 0
	maxIdx := 0
	for idx, val := range arr {
		if val > maxVal {
			maxVal = val
			maxIdx = idx
		}
	}

	return maxIdx, maxVal
}
