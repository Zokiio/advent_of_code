package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var mulPattern = `mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`
var actionPattern = `do\(\)|don't\(\)`

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
	parsedInput := regexp.MustCompile(mulPattern).FindAllStringSubmatch(input, -1)

	var total int
	for _, match := range parsedInput {
		left, err := strconv.Atoi(match[1])
		right, err2 := strconv.Atoi(match[2])
		if err != nil || err2 != nil {
			continue
		}
		total += left * right
	}

	return total
}

func part2(input string) int {
	parsedInput := regexp.MustCompile(mulPattern+"|"+actionPattern).FindAllStringSubmatch(input, -1)
	var total int
	var do bool = true
	for _, match := range parsedInput {
		if match[0] == "do()" {
			do = true
			continue
		}
		if match[0] == "don't()" || !do {
			do = false
			continue
		}
		left, err := strconv.Atoi(match[1])
		right, err2 := strconv.Atoi(match[2])
		if err != nil || err2 != nil {
			continue
		}
		total += left * right
	}

	return total
}
