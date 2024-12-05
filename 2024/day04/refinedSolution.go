package main

func getChar(grid [][]rune, x, y int) (rune, bool) {
	if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		return grid[y][x], true
	}
	return 0, false
}

func points(grid [][]rune) [][2]int {
	var result [][2]int
	for y := range grid {
		for x := range grid[y] {
			result = append(result, [2]int{x, y})
		}
	}
	return result
}

func findXMAS(input string) int {
	grid := parseInput(input)

	directions := []struct {
		dx, dy int
	}{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}

	count := 0

	for _, point := range points(grid) {
		x, y := point[0], point[1]
		char, ok := getChar(grid, x, y)
		if !ok || char != 'X' {
			continue
		}
		for _, dir := range directions {
			matches := true
			expected := []rune{'X', 'M', 'A', 'S'}
			for i := 0; i < 4; i++ {
				char, ok := getChar(grid, x+dir.dx*i, y+dir.dy*i)
				if !ok || char != expected[i] {
					matches = false
					break
				}
			}
			if matches {
				count++
			}
		}
	}

	return count
}

func findMAS(input string) int {
	grid := parseInput(input)

	patterns := []struct {
		points   [][2]int
		expected []rune
	}{
		{
			points:   [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			expected: []rune{'M', 'M', 'S', 'S'},
		},
		{
			points:   [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			expected: []rune{'S', 'S', 'M', 'M'},
		},
		{
			points:   [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			expected: []rune{'M', 'S', 'M', 'S'},
		},
		{
			points:   [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			expected: []rune{'S', 'M', 'S', 'M'},
		},
	}

	count := 0

	for _, point := range points(grid) {
		x, y := point[0], point[1]
		char, ok := getChar(grid, x, y)
		if !ok || char != 'A' {
			continue
		}

		for _, pattern := range patterns {
			matches := true
			for i, offset := range pattern.points {
				char, ok := getChar(grid, x+offset[0], y+offset[1])
				if !ok || char != pattern.expected[i] {
					matches = false
					break
				}
			}
			if matches {
				count++
			}
		}
	}

	return count
}
