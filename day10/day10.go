package day10

import (
	"bufio"
	"os"
)

type point struct {
	x, y int
}

type topoMap [][]int

func (t topoMap) height(x, y int) int {
	if y < 0 || y >= len(t) || x < 0 || x >= len(t[y]) {
		return -1
	}
	return t[y][x]
}

func readFile(filename string) (topoMap, []point) {
	theFile, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer theFile.Close()

	topographicMap := make(topoMap, 0)
	trailheads := make([]point, 0)

	scanner := bufio.NewScanner(theFile)
	scanner.Split(bufio.ScanLines)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for x, h := range line {
			height := int(h - '0')
			row = append(row, height)
			if height == 0 {
				trailheads = append(trailheads, point{x: x, y: y})
			}
		}
		topographicMap = append(topographicMap, row)
	}
	return topographicMap, trailheads
}

func copy(src, target map[point]bool) {
	for k, v := range src {
		target[k] = v
	}
}

func hikeForScore(current point, topographicMap topoMap) map[point]bool {
	currentHeight := topographicMap.height(current.x, current.y)
	if currentHeight == 9 {
		return map[point]bool{current: true}
	}

	foundTops := make(map[point]bool)

	if topographicMap.height(current.x+1, current.y) == currentHeight+1 {
		tmp := hikeForScore(point{x: current.x + 1, y: current.y}, topographicMap)
		copy(tmp, foundTops)
	}

	if topographicMap.height(current.x, current.y+1) == currentHeight+1 {
		tmp := hikeForScore(point{x: current.x, y: current.y + 1}, topographicMap)
		copy(tmp, foundTops)
	}

	if topographicMap.height(current.x-1, current.y) == currentHeight+1 {
		tmp := hikeForScore(point{x: current.x - 1, y: current.y}, topographicMap)
		copy(tmp, foundTops)
	}

	if topographicMap.height(current.x, current.y-1) == currentHeight+1 {
		tmp := hikeForScore(point{x: current.x, y: current.y - 1}, topographicMap)
		copy(tmp, foundTops)
	}

	return foundTops
}

func Part1(filename string) int {
	totalScore := 0
	topographicMap, trailheads := readFile(filename)

	for _, t := range trailheads {
		totalScore += len(hikeForScore(t, topographicMap))
	}

	return totalScore
}

func hikeForRating(current point, topographicMap topoMap) int {
	currentHeight := topographicMap.height(current.x, current.y)
	if currentHeight == 9 {
		return 1
	}

	rating := 0

	if topographicMap.height(current.x+1, current.y) == currentHeight+1 {
		rating += hikeForRating(point{x: current.x + 1, y: current.y}, topographicMap)
	}

	if topographicMap.height(current.x, current.y+1) == currentHeight+1 {
		rating += hikeForRating(point{x: current.x, y: current.y + 1}, topographicMap)
	}

	if topographicMap.height(current.x-1, current.y) == currentHeight+1 {
		rating += hikeForRating(point{x: current.x - 1, y: current.y}, topographicMap)
	}

	if topographicMap.height(current.x, current.y-1) == currentHeight+1 {
		rating += hikeForRating(point{x: current.x, y: current.y - 1}, topographicMap)
	}

	return rating
}

func Part2(filename string) int {
	totalRating := 0
	topographicMap, trailheads := readFile(filename)

	for _, t := range trailheads {
		totalRating += hikeForRating(t, topographicMap)
	}

	return totalRating
}
