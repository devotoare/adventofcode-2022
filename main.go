package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	defer timer("main")()

	readFile, err := os.Open("resources/input.txt")
	if err != nil {
		panic(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	currTotal := 0
	var packs []int

	for _, line := range fileLines {
		if len(line) > 0 {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currTotal = currTotal + value
		} else {
			packs = append(packs, currTotal)
			currTotal = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packs)))

	total := 0

	for i := 0; i < 3; i++ {
		total = total + packs[i]
	}

	fmt.Println(total)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
