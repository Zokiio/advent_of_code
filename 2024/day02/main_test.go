package main

import "testing"

var example = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 2},
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
		{example, 4},
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
