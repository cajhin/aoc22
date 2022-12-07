package main

import (
	"bufio"
	"fmt"
	"os"
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

	//parse. Go array handling sucks a bit. Or rather sucks a kilobyte.
	//translate ABC and XYZ to 123
	var rps1 []byte
	var rps2 []byte

	for _, line := range fileLines {
		if line == "" {
			continue
		}

		barr := []byte(line)
		val1 := 1 + (barr[0] - 'A')
		val2 := 1 + (barr[2] - 'X')
		rps1 = append(rps1, val1)
		rps2 = append(rps2, val2)
		if Testmode {
			fmt.Println(line, val1, val2, "score:", GetScore(val1, val2))
		}
	}

	totalScore := 0
	for i := 0; i < len(rps1); i++ {
		totalScore += GetScore(rps1[i], rps2[i])
	}

	fmt.Println("day2a solution: ", totalScore)
}

func GetScore(a byte, b byte) int {
	score := int(b)
	if a == b {
		score += 3
	} else if a == b-1 || a == b+2 {
		score += 6
	}

	return score
}
