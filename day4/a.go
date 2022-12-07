package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Testmode bool = false
var fileLines []string

func main() {

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Testing")
		Testmode = true
		ReadFileIntoGlobalArray("test")
	} else {
		ReadFileIntoGlobalArray("input")
	}

	//parse, eval
	total := 0
	for _, line := range fileLines {
		pair := strings.Split(line, ",")
		if Testmode {
			fmt.Println(pair[0], "-", pair[1])
		}
		st1, end1 := GetRangeInt(pair[0])
		st2, end2 := GetRangeInt(pair[1])
		if (st1 >= st2 && end1 <= end2) ||
			(st1 <= st2 && end1 >= end2) {
			total++
		}
	}

	fmt.Println("day4a solution: ", total, "is the number of contained ranges")
}

func GetRangeInt(str string) (int, int) {
	tmp := strings.Split(str, "-")
	val1, _ := strconv.Atoi(tmp[0])
	val2, _ := strconv.Atoi(tmp[1])
	return val1, val2
}

func ReadFileIntoGlobalArray(filePath string) {
	//io
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
