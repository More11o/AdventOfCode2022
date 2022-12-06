package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

///
// This doesn't work. This is the 3rd and 'simplest' way i've tried to write this.
// The idea is to create each stack as a slice and then 'push' & 'pop' values
// Data is taken in as a CSV, space delimited, for easy command manipulation.
///

// I have no idea how to parse the begining of the AoC input data for this.
// Unfortutaly this is a manual entry for the starting positions until i'm better.
var s1 = []string{"S", "T", "H", "F", "W", "R"}
var s2 = []string{"S", "G", "D", "Q", "W"}
var s3 = []string{"B", "T", "W"}
var s4 = []string{"D", "R", "W", "T", "N", "Q", "Z", "J"}
var s5 = []string{"F", "B", "H", "G", "L", "V", "T", "Z"}
var s6 = []string{"L", "P", "T", "C", "V", "B", "S", "G"}
var s7 = []string{"Z", "B", "R", "T", "W", "G", "P"}
var s8 = []string{"N", "G", "M", "T", "C", "J", "R"}
var s9 = []string{"L", "G", "B", "W"}

// End of config

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

	data, err := readCSV("input.csv")
	if err != nil {
		panic(err)
	}

	stacks := make(map[int][]string)
	stacks[1] = s1
	stacks[2] = s2
	stacks[3] = s3
	stacks[4] = s4
	stacks[5] = s5
	stacks[6] = s6
	stacks[7] = s7
	stacks[8] = s8

	stacks[9] = s9

	for _, i := range *data {

		qty, _ := strconv.Atoi(i[1])
		unload, _ := strconv.Atoi(i[3])
		load, _ := strconv.Atoi(i[5])

		// Uncomment me to print out the instructions
		//fmt.Printf("move %d from %d to %d \n", qty, unload, load)

		for x := 0; x == qty; x++ {
			l := len(stacks[unload])
			crate := stacks[unload][l]
			stacks[unload] = stacks[unload][:l-1]

			stacks[load] = append(stacks[load], crate)
		}
	}

	for x := 0; x < 9; x++ {
		fmt.Println(stacks[x])
	}

}
