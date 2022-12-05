package main

import (
	"bufio"
	"fmt"
	"os"
)

func editPriority(s []rune) []rune {
	for i, char := range s {
		// unicode a < A
		if char < 'a' {
			s[i] = char - 38 // magic number
		} else {
			s[i] = char - 96 // magic number
		}
	}

	return s
}

func findDuplicates(f *[]rune, s *[]rune) rune {
	for _, fchar := range *f {
		for _, schar := range *s {
			if fchar == schar {
				return fchar
			}
		}
	}

	return ' '
}

func findBadge(s1 string, s2 string, s3 string) rune {

	for _, elf1item := range s1 {
		for _, elf2item := range s2 {
			for _, elf3item := range s3 {
				if (elf1item == elf2item) && (elf1item == elf3item) {
					return elf1item
				}
			}
		}
	}
	return ' '
}

func readfile(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data := bufio.NewScanner(file)
	data.Split(bufio.ScanLines)

	for data.Scan() {
		lines = append(lines, data.Text())
	}

	return lines
}

func main() {

	var sumPriorities int = 0

	data := readfile("input.txt")

	for _, rucksack := range data {

		items := []rune(rucksack)

		items = editPriority(items)

		firstCompartment := items[:len(items)/2]
		secondCompartment := items[len(items)/2:]

		duplicate := findDuplicates(&firstCompartment, &secondCompartment)

		sumPriorities += int(duplicate)
	}

	// Answer 1
	//fmt.Println(sumPriorities)

	var badges []rune
	var answer int = 0
	for i := 0; i < len(data); i += 3 {
		badges = append(badges, findBadge(data[i], data[i+1], data[i+2]))
	}
	badges = editPriority(badges)

	for _, badge := range badges {
		answer += int(badge)
	}

	fmt.Println(answer)
}
