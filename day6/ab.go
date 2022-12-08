package main

import (
	"bufio"
	"fmt"
	"os"
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

	if Testmode {
		Test()
	}

	line := fileLines[0]
	posA := -1
	for i := 4; i < len(line); i++ {
		if !HasDuplicateChars(line[i-4 : i]) {
			posA = i
			break
		}
	}

	posB := -1
	for i := 14; i < len(line); i++ {
		if !HasDuplicateChars(line[i-14 : i]) {
			posB = i
			break
		}
	}

	fmt.Println("day6a solution: first 4 unique chars end at pos:", posA)
	fmt.Println("day6b solution: first 14 unique chars end at pos:", posB)
}

func HasDuplicateChars(str string) bool {
	for i := len(str) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if str[i] == str[j] {
				return true
			}
		}
	}

	return false
}

func Test() {
	if HasDuplicateChars("abcdABCD") || HasDuplicateChars("") || HasDuplicateChars(("X")) ||
		!HasDuplicateChars("abcdeabcde") || !HasDuplicateChars("abcdebcdea") {
		fmt.Println("Test FAILED")
		os.Exit(0)
	}

}

func ReadFileIntoGlobalArray(filePath string) {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		l := fileScanner.Text()
		if len(l) > 0 && l[0] != '#' {
			fileLines = append(fileLines, l)
		} else if Testmode {
			fmt.Println(l)
		}
	}
	readFile.Close()
}
