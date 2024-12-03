package main

import "testing"

var example = `
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`

func Test_part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example, 161},
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

var example2 = `
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
`

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{example2, 48},
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
