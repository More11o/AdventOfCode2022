package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	id       int
	calories int
}

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var lines []string
	var elves []Elf

	file, err := os.Open("input.txt")
	check_error(err)

	data := bufio.NewScanner(file)
	data.Split(bufio.ScanLines)

	for data.Scan() {
		lines = append(lines, data.Text())
	}

	// Getting a . weird error with the slice.
	// elves before this append has a length of zero and no object to assign data to
	elves = append(elves, Elf{id: 0, calories: 0})
	// At this point elves[] len is 1 but the struct is at index 0

	for _, line := range lines {
		if line == "" {
			elves = append(elves, Elf{id: len(elves), calories: 0})
		} else {
			calories, err := strconv.Atoi(line)
			check_error(err)
			elves[len(elves)-1].calories += calories
		}

	}

	//Sort the structs by the calories.
	// This allows index 0 to be the greatest amount but maintains Elf.id to identify which elf has that many calories.
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})

	// Answer 01
	//fmt.Println(elves[0])

	//Answer 02
	fmt.Println(elves[0].calories + elves[1].calories + elves[2].calories)

}
