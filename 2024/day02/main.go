package main

import (
	_ "embed"
	"fmt"
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
	levels := trimSplit(input)
	var totalSafe int

	for _, level := range levels {
		if isSafeSequence(level) {
			totalSafe++
		}
	}
	return totalSafe
}

func part2(input string) int {
	levels := trimSplit(input)
	var totalSafe int

	for _, level := range levels {
		if len(level) < 2 {
			continue
		}

		if isSafeSequence(level) {
			totalSafe++
			continue
		}

		for i := 0; i < len(level); i++ {
			newLevel := make([]int, 0, len(level)-1)
			newLevel = append(newLevel, level[:i]...)
			newLevel = append(newLevel, level[i+1:]...)

			if isSafeSequence(newLevel) {
				totalSafe++
				break
			}
		}
	}
	return totalSafe
}

func isSafeSequence(level []int) bool {
	if len(level) < 2 {
		return false
	}

	prevNum := level[0]
	diff := level[1] - level[0]
	if diff == 0 || absDiff(diff) > 3 {
		return false
	}

	isIncreasing := diff > 0

	for i := 1; i < len(level); i++ {
		num := level[i]
		diff = num - prevNum

		if diff == 0 || absDiff(diff) > 3 {
			return false
		}

		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}

		prevNum = num
	}

	return true
}

func absDiff(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func trimSplit(input string) [][]int {
	var levels [][]int

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var level []int
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			level = append(level, num)
		}
		levels = append(levels, level)
	}
	return levels
}
