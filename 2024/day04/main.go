package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("input is empty")
	}
}

func main() {
	part1Answer := part1(input)
	fmt.Println("Output:", part1Answer)

	part2Answer := part2(input)
	fmt.Println("Output:", part2Answer)
}

func part1(input string) int {
	countXmas := 0
	var inputRunes [][]rune

	splitString := strings.Split(strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n")), "\n")
	for _, s := range splitString {
		inputRunes = append(inputRunes, []rune(s))
	}

	for i, row := range inputRunes {
		for j := 0; j < len(row); j++ {
			// hor
			if j <= (len(row) - 4) {
				if row[j] == 'X' && row[j+1] == 'M' && row[j+2] == 'A' && row[j+3] == 'S' {
					countXmas++
				}
			}
			// rev hor
			if j >= 3 {
				if row[j] == 'X' && row[j-1] == 'M' && row[j-2] == 'A' && row[j-3] == 'S' {
					countXmas++
				}
			}
			// up
			if i >= 3 {
				if row[j] == 'X' && inputRunes[i-1][j] == 'M' && inputRunes[i-2][j] == 'A' && inputRunes[i-3][j] == 'S' {
					countXmas++
				}
			}
			// diagonal upwards right
			if i >= 3 && j <= (len(row)-4) {
				if row[j] == 'X' && inputRunes[i-1][j+1] == 'M' && inputRunes[i-2][j+2] == 'A' && inputRunes[i-3][j+3] == 'S' {
					countXmas++
				}
			}
			// diagonal upwards left
			if i >= 3 && j >= 3 {
				if row[j] == 'X' && inputRunes[i-1][j-1] == 'M' && inputRunes[i-2][j-2] == 'A' && inputRunes[i-3][j-3] == 'S' {
					countXmas++
				}
			}

			// down
			if i <= (len(inputRunes) - 4) {
				if row[j] == 'X' && inputRunes[i+1][j] == 'M' && inputRunes[i+2][j] == 'A' && inputRunes[i+3][j] == 'S' {
					countXmas++
				}
			}
			// diagonal downwards right
			if i <= (len(inputRunes)-4) && j <= (len(row)-4) {
				if row[j] == 'X' && inputRunes[i+1][j+1] == 'M' && inputRunes[i+2][j+2] == 'A' && inputRunes[i+3][j+3] == 'S' {
					countXmas++
				}
			}
			// diagonal downwards left
			if i <= (len(inputRunes)-4) && j >= 3 {
				if row[j] == 'X' && inputRunes[i+1][j-1] == 'M' && inputRunes[i+2][j-2] == 'A' && inputRunes[i+3][j-3] == 'S' {
					countXmas++
				}
			}
		}
	}
	return countXmas
}

func part2(input string) int {
	countXmas := 0
	var inputRunes [][]rune

	splitString := strings.Split(strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n")), "\n")
	for _, s := range splitString {
		inputRunes = append(inputRunes, []rune(s))
	}

	for i, row := range inputRunes {
		if i < 1 || i >= len(inputRunes)-1 {
			continue
		}
		for j := 0; j < len(row); j++ {
			if j < 1 || j > len(row)-2 {
				continue
			}

			if row[j] != 'A' {
				continue
			}

			if inputRunes[i-1][j-1] == 'M' && inputRunes[i-1][j+1] == 'M' &&
				inputRunes[i+1][j-1] == 'S' && inputRunes[i+1][j+1] == 'S' {
				countXmas++
			}

			if inputRunes[i-1][j-1] == 'S' && inputRunes[i-1][j+1] == 'S' &&
				inputRunes[i+1][j-1] == 'M' && inputRunes[i+1][j+1] == 'M' {
				countXmas++
			}

			if inputRunes[i-1][j-1] == 'M' && inputRunes[i-1][j+1] == 'S' &&
				inputRunes[i+1][j-1] == 'M' && inputRunes[i+1][j+1] == 'S' {
				countXmas++
			}

			if inputRunes[i-1][j-1] == 'S' && inputRunes[i-1][j+1] == 'M' &&
				inputRunes[i+1][j-1] == 'S' && inputRunes[i+1][j+1] == 'M' {
				countXmas++
			}
		}
	}

	return countXmas
}
