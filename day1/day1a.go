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

	//count
	calories := []int{0}
	elfidx := 0
	maxelf := 0
	for _, line := range fileLines {
		if line == "" {
			if calories[elfidx] > calories[maxelf] {
				maxelf = elfidx
			}
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

	fmt.Println("Gold medal goes to elf #", maxelf+1, "of", len(calories)+1, "with", calories[maxelf], "calories")
}
