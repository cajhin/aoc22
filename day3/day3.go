package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Testmode bool = false

func main() {

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Testing")
		Testmode = true
	}

	//io
	filePath := "input"
	if Testmode {
		filePath = "test"
	}
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
	total := 0
	for i := 0; i+2 <= len(fileLines); i += 3 {
		dup := FindDuplicate(fileLines[i], fileLines[i+1], fileLines[i+2])
		total += RuneToPrio(dup)
		if Testmode {
			fmt.Println(i, "-", "dup:", string(dup))
		}
	}

	fmt.Println("day3b solution: ", total, "is the added priority of all badge runes")
}

func FindDuplicate(a string, b string, c string) rune {

	for _, ru := range a {
		if strings.Contains(b, string(ru)) &&
			strings.Contains(c, string(ru)) {
			return ru
		}
	}

	return -1
}

func RuneToPrio(ru rune) int {
	if ru >= 'a' {
		return 1 + int(ru-'a')
	}
	return 27 + int(ru-'A')
}
