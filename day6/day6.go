package day6

import (
	"bufio"
	"os"
)

type movement func(p point) point

var movements = []movement{
	func(p point) point { return point{x: p.x, y: p.y - 1} },
	func(p point) point { return point{x: p.x + 1, y: p.y} },
	func(p point) point { return point{x: p.x, y: p.y + 1} },
	func(p point) point { return point{x: p.x - 1, y: p.y} },
}

type area [][]rune

func (a area) inBounds(p point) bool {
	return p.y >= 0 && p.y < len(a) && p.x >= 0 && p.x < len(a[p.y])
}

func (a area) obstacle(p point) bool {
	if a.inBounds(p) {
		return a[p.y][p.x] == '#'
	}
	return false
}

func (a area) space() int {
	return len(a) * len(a[0])
}

type point struct {
	x, y int
}

func readFile(filename string) (area, point) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	a := make(area, 0)
	start := point{}

	y := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0, len(line))
		for x, c := range line {
			row = append(row, c)
			if c == '^' {
				start.x, start.y = x, y
			}
		}
		a = append(a, row)
		y++
	}

	return a, start
}

func Part1(filename string) int {
	visited := make(map[point]bool)
	a, pos := readFile(filename)
	moveIndex := 0

	for a.inBounds(pos) {
		visited[pos] = true
		newPos := movements[moveIndex](pos)
		if a.obstacle(newPos) {
			moveIndex = (moveIndex + 1) % 4
			newPos = movements[moveIndex](pos)
		}
		pos = newPos
	}

	return len(visited)
}

func loopCount(guardRoute map[point]bool, start point, a area) int {
	loops := 0
	for p := range guardRoute {
		if p != start {
			a[p.y][p.x] = '#'

			currentPos := start
			moveIndex := 0
			steps := 0
			for a.inBounds(currentPos) && steps <= a.space() {
				steps++
				newPos := movements[moveIndex](currentPos)
				for a.obstacle(newPos) {
					moveIndex = (moveIndex + 1) % 4
					newPos = movements[moveIndex](currentPos)
				}
				currentPos = newPos
			}
			a[p.y][p.x] = '.'

			if a.inBounds(currentPos) {
				loops++
			}
		}
	}
	return loops
}

func Part2(filename string) int {
	visited := make(map[point]bool)
	a, start := readFile(filename)
	moveIndex := 0

	currentPos := start
	for a.inBounds(currentPos) {
		visited[currentPos] = true
		newPos := movements[moveIndex](currentPos)
		for a.obstacle(newPos) {
			moveIndex = (moveIndex + 1) % 4
			newPos = movements[moveIndex](currentPos)
		}
		currentPos = newPos
	}

	return loopCount(visited, start, a)
}
