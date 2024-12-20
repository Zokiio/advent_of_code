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

	// improvedAnswer
	part1Refined := findXMAS(input)
	fmt.Println("refined Output:", part1Refined)

	part2Refined := findMAS(input)
	fmt.Println("refined Output:", part2Refined)

}

func parseInput(input string) [][]rune {
	var inputRunes [][]rune
	splitString := strings.Split(strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n")), "\n")
	for _, s := range splitString {
		inputRunes = append(inputRunes, []rune(s))
	}
	return inputRunes
}

func part1(input string) int {
	countXmas := 0
	inputRunes := parseInput(input)

	for i, row := range inputRunes {
		for j := 0; j < len(row); j++ {
			// ➡️
			if j <= (len(row) - 4) {
				if row[j] == 'X' && row[j+1] == 'M' && row[j+2] == 'A' && row[j+3] == 'S' {
					countXmas++
				}
			}
			// ⬅️
			if j >= 3 {
				if row[j] == 'X' && row[j-1] == 'M' && row[j-2] == 'A' && row[j-3] == 'S' {
					countXmas++
				}
			}
			// ⬆️
			if i >= 3 {
				if row[j] == 'X' && inputRunes[i-1][j] == 'M' && inputRunes[i-2][j] == 'A' && inputRunes[i-3][j] == 'S' {
					countXmas++
				}
			}
			// ↗️
			if i >= 3 && j <= (len(row)-4) {
				if row[j] == 'X' && inputRunes[i-1][j+1] == 'M' && inputRunes[i-2][j+2] == 'A' && inputRunes[i-3][j+3] == 'S' {
					countXmas++
				}
			}
			// ↖️
			if i >= 3 && j >= 3 {
				if row[j] == 'X' && inputRunes[i-1][j-1] == 'M' && inputRunes[i-2][j-2] == 'A' && inputRunes[i-3][j-3] == 'S' {
					countXmas++
				}
			}

			// ⬇️
			if i <= (len(inputRunes) - 4) {
				if row[j] == 'X' && inputRunes[i+1][j] == 'M' && inputRunes[i+2][j] == 'A' && inputRunes[i+3][j] == 'S' {
					countXmas++
				}
			}
			// ↘️
			if i <= (len(inputRunes)-4) && j <= (len(row)-4) {
				if row[j] == 'X' && inputRunes[i+1][j+1] == 'M' && inputRunes[i+2][j+2] == 'A' && inputRunes[i+3][j+3] == 'S' {
					countXmas++
				}
			}
			// ↙️
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
	inputRunes := parseInput(input)

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
