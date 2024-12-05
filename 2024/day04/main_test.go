package main

import "testing"

var example = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 18},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("part1() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 9},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("part2() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_findXMAS(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 18},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := findXMAS(test.input)
			if got != test.want {
				t.Errorf("findXMAS() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_findMAS(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 9},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := findMAS(test.input)
			if got != test.want {
				t.Errorf("findMAS() = %v, want %v", got, test.want)
			}
		})
	}
}
