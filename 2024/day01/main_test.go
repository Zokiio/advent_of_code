package main

import "testing"

var example = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 11},
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
		{example, 31},
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
