package main

import (
	"fmt"
	"os"
)

// I get completly replaced by answer02
func answer01(buf []byte) int {
	for x := 0; x < len(buf)-4; x++ {

		// build a map with the buffer byte as the key and an empty struct as the value.
		// empty structs don't contain any space but allows the use of map functions on the keys.
		bytes := make(map[byte]struct{})
		bytes[buf[x]] = struct{}{}
		bytes[buf[x+1]] = struct{}{}
		bytes[buf[x+2]] = struct{}{}
		bytes[buf[x+3]] = struct{}{}

		// this will only return unique keys.
		// if it's less than 4 in length then they can't have been unique.
		var check []byte
		for b := range bytes {
			check = append(check, b)
		}
		if len(check) == 4 {
			// answer wants the index of the last byte
			return x + 4
		}

	}
	return 0
}

func answer02(buf []byte, size int) int {
	for x := 0; x < len(buf); x++ {

		// build a map with the buffer byte as the key and an empty struct as the value.
		// empty structs don't contain any space but allows the use of map functions on the keys.
		// buffer is of length size
		bytes := make(map[byte]struct{})
		for y := 0; y < size; y++ {
			bytes[buf[x+y]] = struct{}{}
		}

		// this will only return unique keys.
		// if it's less than 'size' in length then they can't have been unique.
		var check []byte
		for b := range bytes {
			check = append(check, b)
		}
		if len(check) == size {

			// answers want the index in relation to the size
			return x + size
		}

	}

	return 0
}

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(answer02(buf, 4))
	fmt.Println(answer02(buf, 14))

}
