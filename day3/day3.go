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
	var comp1 []string
	var comp2 []string

	for _, line := range fileLines {
		if line == "" {
			continue
		}
		len := len(line)
		str1 := line[0 : len/2]
		str2 := line[len/2 : len]

		comp1 = append(comp1, str1)
		comp2 = append(comp2, str2)
		if Testmode {
			fmt.Println(line, str1, str2)
		}
	}

	total := 0
	for i := 0; i < len(comp1); i++ {
		dup := FindDuplicate(comp1[i], comp2[i])
		prio := RuneToPrio(dup)
		total += prio
		if Testmode {
			fmt.Println("dup: ", string(dup), "=", prio)
		}
	}

	fmt.Println("day3a solution: ", total, "is the added priority of all duplicate runes")
}

func FindDuplicate(a string, b string) rune {

	for _, ru := range a {
		if strings.Contains(b, string(ru)) {
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
