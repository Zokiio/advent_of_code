package main

import (
	_ "embed"
	"testing"
)

//go:embed example_input.txt
var example string

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 143},
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
		{example, 123},
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
