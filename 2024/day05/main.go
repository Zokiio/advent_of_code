package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n"))
	if len(input) == 0 {
		panic("input is empty")
	}
}

func main() {
	part1Answer := part1(input)
	fmt.Println("Part 1:", part1Answer)

	part2Answer := part2(input)
	fmt.Println("Part 2:", part2Answer)
}

type SafetyManual struct {
	rules   map[string][]string
	updates [][]string
}

func parseInput(input string) SafetyManual {
	rules := map[string][]string{}
	updates := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "|") {
			v := strings.Split(line, "|")
			if len(v) == 2 {
				key := strings.TrimSpace(v[0])
				value := strings.TrimSpace(v[1])
				rules[key] = append(rules[key], value)
			}
		} else if line != "" {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return SafetyManual{rules: rules, updates: updates}
}

func isValidOrder(update []string, rules map[string][]string) bool {
	for i := 0; i < len(update); i++ {
		number := update[i]
		for _, rule := range rules[number] {
			previous_nums := update[:i]
			if slices.Index(previous_nums, rule) != -1 {
				return false
			}
		}
	}
	return true
}

func part1(input string) int {
	manual := parseInput(input)
	count := 0

	for _, update := range manual.updates {
		if isValidOrder(update, manual.rules) {
			val, _ := strconv.Atoi(update[len(update)/2])
			count += val
		}
	}

	return count
}

func part2(input string) int {
	manual := parseInput(input)
	count := 0

	for _, update := range manual.updates {
		currentUpdate := slices.Clone(update)
		valid := true
		for i := 0; i < len(currentUpdate); i++ {
			number := currentUpdate[i]
			for _, rule := range manual.rules[number] {
				previous_nums := currentUpdate[:i]
				index := slices.Index(previous_nums, rule)
				if index != -1 {
					currentUpdate = slices.Delete(currentUpdate, index, index+1)
					currentUpdate = slices.Insert(currentUpdate, i, rule)
					valid = false
					i = 0
					break
				}
			}
		}
		if !valid {
			val, _ := strconv.Atoi(currentUpdate[len(currentUpdate)/2])
			count += val
		}
	}

	return count
}
