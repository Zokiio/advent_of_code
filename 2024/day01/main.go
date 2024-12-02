package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
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
	lines := trimSplit(input)
	var first, second []int

	for _, line := range lines {
		splitStr := strings.Fields(line)
		if len(splitStr) != 2 {
			continue
		}

		fInt, err := strconv.Atoi(splitStr[0])
		sInt, err1 := strconv.Atoi(splitStr[1])
		if err != nil || err1 != nil {
			continue
		}

		first = append(first, fInt)
		second = append(second, sInt)
	}

	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})

	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})

	sum := 0
	for i := range first {
		diff := first[i] - second[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}

func part2(input string) int {
	lines := trimSplit(input)

	var first, second []int

	for _, line := range lines {
		splitStr := strings.Fields(line)
		if len(splitStr) != 2 {
			continue
		}

		fInt, err := strconv.Atoi(splitStr[0])
		sInt, err1 := strconv.Atoi(splitStr[1])
		if err != nil || err1 != nil {
			continue
		}

		first = append(first, fInt)
		second = append(second, sInt)
	}

	var total int
	for _, first := range first {
		var times int
		for _, second := range second {
			if first == second {
				times++
			}
		}
		total += first * times
	}
	return total
}

func trimSplit(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	return lines
}
