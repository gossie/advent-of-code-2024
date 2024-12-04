package day4

import (
	"bufio"
	"os"
)

func readFile(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	field := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		field = append(field, row)
	}

	return field
}

func search(field [][]rune, y, x int) int {
	findings := 0

	// horizontally to the right
	chars := make([]rune, 0, 4)
	for i := x; i < len(field[y]) && i < x+4; i++ {
		chars = append(chars, field[y][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// first diagonal
	chars = make([]rune, 0, 4)
	for i, j := x, y; i < len(field[y]) && i < x+4 && j < len(field) && j < y+4; i, j = i+1, j+1 {
		chars = append(chars, field[j][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// vertically down
	chars = make([]rune, 0, 4)
	for i := y; i < len(field) && i < y+4; i++ {
		chars = append(chars, field[i][x])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// second diagonal
	chars = make([]rune, 0, 4)
	for i, j := x, y; i >= 0 && i > x-4 && j < len(field) && j < y+4; i, j = i-1, j+1 {
		chars = append(chars, field[j][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// horizontally to the left
	chars = make([]rune, 0, 4)
	for i := x; i >= 0 && i > x-4; i-- {
		chars = append(chars, field[y][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// third diagonal
	chars = make([]rune, 0, 4)
	for i, j := x, y; i >= 0 && i > x-4 && j >= 0 && j > y-4; i, j = i-1, j-1 {
		chars = append(chars, field[j][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// vertically up
	chars = make([]rune, 0, 4)
	for i := y; i >= 0 && i > y-4; i-- {
		chars = append(chars, field[i][x])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	// fourth diagonal
	chars = make([]rune, 0, 4)
	for i, j := x, y; i < len(field[y]) && i < x+4 && j >= 0 && j > y-4; i, j = i+1, j-1 {
		chars = append(chars, field[j][i])
	}
	if string(chars) == "XMAS" {
		findings++
	}

	return findings
}

func check(field [][]rune, y, x int) bool {
	return (field[y-1][x-1] == 'M' && field[y+1][x+1] == 'S' || field[y-1][x-1] == 'S' && field[y+1][x+1] == 'M') && (field[y-1][x+1] == 'M' && field[y+1][x-1] == 'S' || field[y-1][x+1] == 'S' && field[y+1][x-1] == 'M')
}

func Part1(filename string) int {
	count := 0
	field := readFile(filename)

	for y := range len(field) {
		for x := range len(field[y]) {
			if field[y][x] == 'X' {
				count += search(field, y, x)
			}
		}
	}

	return count
}

func Part2(filename string) int {
	count := 0
	field := readFile(filename)

	for y := 1; y < len(field)-1; y++ {
		for x := 1; x < len(field[y])-1; x++ {
			if field[y][x] == 'A' {
				if check(field, y, x) {
					count++
				}
			}
		}
	}

	return count
}
