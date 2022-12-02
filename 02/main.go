package main

import (
	"encoding/csv"
	"os"
)

// FIXME:
// Theres a way to make this type of solution work but i haven't figured it out yet.
func play_part_one_bug(o int, p int) int {
	// Draw
	if o == p {
		return 3 + p
	}
	// Win
	if (o < p) || (o-p == 2) {
		return 6 + p
	}

	// Loose
	if (o > p) || (o-p == -2) {
		return p
	}

	// Should really return an error and deal with it gracefully but i don't want to check it every loop.
	// if it errors here the program is dead regardless.
	panic("failed to compute round")
}

// This is an ugly way of doing it.
func play_part_one(o int, p int) int {
	switch o {
	case 1:
		switch p {
		case 1:
			return p + 3
		case 2:
			return p + 6
		case 3:
			return p
		}
	case 2:
		switch p {
		case 1:
			return p
		case 2:
			return p + 3
		case 3:
			return p + 6
		}
	case 3:
		switch p {
		case 1:
			return p + 6
		case 2:
			return p
		case 3:
			return p + 3
		}

	}

	// Should really return an error and deal with it gracefully but i don't want to check it every loop.
	// if it errors here the program is dead regardless.
	panic("failed to compute round")
}

func play_part_two(o int, p int) int {
	switch o {
	case 1:
		switch p {
		case 1: //loose
			return 3
		case 2: // draw
			return 1 + 3
		case 3: //win
			return 2 + 6
		}
	case 2:
		switch p {
		case 1:
			return 1
		case 2:
			return 2 + 3
		case 3:
			return 3 + 6
		}
	case 3:
		switch p {
		case 1:
			return 2
		case 2:
			return 3 + 3
		case 3:
			return 1 + 6
		}

	}

	// Should really return an error and deal with it gracefully but i don't want to check it every loop.
	// if it errors here the program is dead regardless.
	panic("failed to compute round")
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
	csvReader.Comma = ' '
	data, err = csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	var score int = 0

	scores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	data, err := readCSV("input.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range *data {
		score += play_part_two(
			scores[line[0]],
			scores[line[1]])
	}

	println(score)
}
