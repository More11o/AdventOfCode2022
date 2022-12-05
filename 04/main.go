package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer01(d *[][]string) int {
	elves := [4]int{0, 0, 0, 0}
	answer := 0
	for _, assignmens := range *d {
		for x := 0; x < len(assignmens); x += 2 {

			str1 := strings.Split(assignmens[x], "-")
			str2 := strings.Split(assignmens[x+1], "-")

			elves[0], _ = strconv.Atoi(str1[0])

			elves[1], _ = strconv.Atoi(str1[1])

			elves[2], _ = strconv.Atoi(str2[0])

			elves[3], _ = strconv.Atoi(str2[1])

		}

		if (elves[0] == elves[2]) && (elves[1] == elves[3]) {
			// Elf 1	Elf 2
			// 1-5		1-5
			// This causes double accounting.
			answer++
			continue
		}

		if (elves[0] <= elves[2]) && (elves[1] >= elves[3]) {
			answer++
		}

		if (elves[0] >= elves[2]) && (elves[1] <= elves[3]) {
			answer++
		}

	}

	return answer
}

func answer02(d *[][]string) int {
	answer := 0
	var elf1 []int
	var elf2 []int
	var overlap bool = false

	for _, assignmens := range *d {
		for x := 0; x < len(assignmens); x += 2 {

			str1 := strings.Split(assignmens[x], "-")
			str2 := strings.Split(assignmens[x+1], "-")

			e1s, _ := strconv.Atoi(str1[0])

			e1e, _ := strconv.Atoi(str1[1])

			e2s, _ := strconv.Atoi(str2[0])

			e2e, _ := strconv.Atoi(str2[1])

			elf1 = buildArray(e1s, e1e)
			elf2 = buildArray(e2s, e2e)

			for _, v := range elf1 {
				if contains(elf2, v) {
					overlap = true
				}
			}

			for _, v := range elf2 {
				if contains(elf1, v) {
					overlap = true
				}
			}

			if overlap {
				answer++
				overlap = false
			}

		}
	}

	return answer
}

func contains[T comparable](r []T, e T) bool {
	for _, v := range r {
		if v == e {
			return true
		}
	}
	return false
}

// e is endpoint inclusive
func buildArray(s int, e int) []int {
	var a []int
	for x := s; x <= e; x++ {
		a = append(a, x)
	}
	return a
}

func readCSV(filename string) (*[][]string, error) {
	var err error
	var data [][]string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	var a int
	data, err := readCSV("input.csv")
	if err != nil {
		panic(err)
	}

	a = answer01(data)
	fmt.Println(a)

	a = answer02(data)
	fmt.Println(a)

}
