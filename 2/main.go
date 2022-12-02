package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// A = Rock - 1
// B = Paper - 2
// C = Scissors - 3
// Loss = 0
// Draw = 3
// Win  = 6
// X = lose
// Y = draw
// Z = win

func main() {
	defer timer("main")()

	readFile, err := os.Open("resources/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	totalScore := 0

	winning := make(map[string]int)
	winning["A"] = 2
	winning["B"] = 3
	winning["C"] = 1

	losing := make(map[string]int)
	losing["A"] = 3
	losing["B"] = 1
	losing["C"] = 2

	drawing := make(map[string]int)
	drawing["A"] = 1
	drawing["B"] = 2
	drawing["C"] = 3

	for _, line := range fileLines {
		thisGame := strings.Split(line, " ")
		totalScore += processPartTwo(thisGame[0], thisGame[1], winning, losing, drawing)
	}

	fmt.Println(totalScore)
}

func processPartTwo(oppPlay string, myPlay string, winning map[string]int, losing map[string]int, drawing map[string]int) int {

	if myPlay == "X" {
		return losing[oppPlay]
	} else if myPlay == "Y" {
		return drawing[oppPlay] + 3
	} else if myPlay == "Z" {
		return winning[oppPlay] + 6
	}

	return 0
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
