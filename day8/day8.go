package day8

import (
	"bufio"
	"math"
	"os"
)

type point struct {
	x, y int
}

type antenna struct {
	frequency rune
	position  point
}

func readFile(filename string) ([]*antenna, map[rune][]*antenna, int, int) {
	theFile, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer theFile.Close()

	antennas := make([]*antenna, 0)
	byFrequency := make(map[rune][]*antenna)
	maxX, maxY := 0, 0

	scanner := bufio.NewScanner(theFile)
	scanner.Split(bufio.ScanLines)
	for y := 0; scanner.Scan(); y++ {
		maxY = int(math.Max(float64(y), float64(maxY)))
		for x, c := range scanner.Text() {
			maxX = int(math.Max(float64(x), float64(maxX)))
			if c != '.' {
				newAntenna := &antenna{frequency: c, position: point{x: x, y: y}}
				antennas = append(antennas, newAntenna)
				_, found := byFrequency[c]
				if found {
					byFrequency[c] = append(byFrequency[c], newAntenna)
				} else {
					byFrequency[c] = []*antenna{newAntenna}
				}
			}
		}

	}

	return antennas, byFrequency, maxX, maxY
}

func Part1(filename string) int {
	antennas, byFrequency, maxX, maxY := readFile(filename)
	positions := make(map[point]bool)
	for _, a := range antennas {
		for _, other := range byFrequency[a.frequency] {
			if a.position != other.position {
				point1 := point{x: a.position.x + (a.position.x - other.position.x), y: a.position.y + (a.position.y - other.position.y)}
				if point1.x >= 0 && point1.x <= maxX && point1.y >= 0 && point1.y <= maxY {
					positions[point1] = true
				}

				point2 := point{x: other.position.x + (other.position.x - a.position.x), y: other.position.y + (other.position.y - a.position.y)}
				if point2.x >= 0 && point2.x <= maxX && point2.y >= 0 && point2.y <= maxY {
					positions[point2] = true
				}
			}
		}
	}
	return len(positions)
}

func Part2(filename string) int {
	antennas, byFrequency, maxX, maxY := readFile(filename)
	positions := make(map[point]bool)
	for _, a := range antennas {
		for _, other := range byFrequency[a.frequency] {
			if a.position != other.position {
				point1 := a.position
				deltaX1 := point1.x - other.position.x
				deltaY1 := point1.y - other.position.y
				for point1.x >= 0 && point1.x <= maxX && point1.y >= 0 && point1.y <= maxY {
					positions[point1] = true
					point1 = point{x: point1.x + deltaX1, y: point1.y + deltaY1}
				}

				point2 := other.position
				deltaX2 := point2.x - a.position.x
				deltaY2 := point2.y - a.position.y
				for point2.x >= 0 && point2.x <= maxX && point2.y >= 0 && point2.y <= maxY {
					positions[point2] = true
					point2 = point{x: point2.x + deltaX2, y: point2.y + deltaY2}
				}
			}
		}
	}
	return len(positions)
}
