package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	line := fileLines[0]
	posA := FindFirstMarker(line, 4)
	posB := FindFirstMarker(line, 14)

	fmt.Println("day6a solution: first 4 unique chars end at pos:", posA)
	fmt.Println("day6b solution: first 14 unique chars end at pos:", posB)
}

func FindFirstMarker(line string, windowSize int) int {
	for i := windowSize; i < len(line); i++ {
		if !HasDuplicateChars(line[i-windowSize : i]) {
			return i
		}
	}

	return -1
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

func jot_day06OnesCount(s string, size int) int {
	for i := 0; i < len(s)-size; i++ {
		var marker uint32
		for j := i; j < i+size; j++ {
			marker |= 1 << (s[j] - 'a')
		}
		if bits.OnesCount32(marker) == size {
			return i + size
		}
	}
	return 0
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
