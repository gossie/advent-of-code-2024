package day12

import (
	"bufio"
	"os"
)

type plot struct {
	plant      rune
	inRegion   bool
	outerSides int
	x, y       int
}

func readFile(filename string) [][]*plot {
	theFile, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer theFile.Close()

	garden := make([][]*plot, 0)

	scanner := bufio.NewScanner(theFile)
	scanner.Split(bufio.ScanLines)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]*plot, 0, len(line))
		for x, p := range line {
			row = append(row, &plot{
				plant:      p,
				inRegion:   false,
				outerSides: 0,
				x:          x,
				y:          y,
			})
		}
		garden = append(garden, row)
		y++
	}

	return garden
}

func scanRegion(garden [][]*plot, x, y int) []*plot {
	region := make([]*plot, 0)
	toHandle := []*plot{garden[y][x]}
	for len(toHandle) > 0 {
		p := toHandle[0]
		toHandle = toHandle[1:]
		if p.inRegion {
			continue
		}

		region = append(region, p)
		p.inRegion = true

		if p.x-1 >= 0 && !garden[p.y][p.x-1].inRegion && garden[p.y][p.x-1].plant == p.plant {
			toHandle = append(toHandle, garden[p.y][p.x-1])
		}
		if p.x-1 < 0 || garden[p.y][p.x-1].plant != p.plant {
			p.outerSides++
		}

		if p.y-1 >= 0 && !garden[p.y-1][p.x].inRegion && garden[p.y-1][p.x].plant == p.plant {
			toHandle = append(toHandle, garden[p.y-1][p.x])
		}
		if p.y-1 < 0 || garden[p.y-1][p.x].plant != p.plant {
			p.outerSides++
		}

		if p.x+1 < len(garden[p.y]) && !garden[p.y][p.x+1].inRegion && garden[p.y][p.x+1].plant == p.plant {
			toHandle = append(toHandle, garden[p.y][p.x+1])
		}
		if p.x+1 >= len(garden[p.y]) || garden[p.y][p.x+1].plant != p.plant {
			p.outerSides++
		}

		if p.y+1 < len(garden) && !garden[p.y+1][p.x].inRegion && garden[p.y+1][p.x].plant == p.plant {
			toHandle = append(toHandle, garden[p.y+1][p.x])
		}
		if p.y+1 >= len(garden) || garden[p.y+1][p.x].plant != p.plant {
			p.outerSides++
		}
	}
	return region
}

func scan(garden [][]*plot) map[int][]*plot {
	nextRegionId := 1
	regions := make(map[int][]*plot)
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[y]); x++ {
			if garden[y][x].inRegion {
				continue
			}

			regions[nextRegionId] = scanRegion(garden, x, y)
			nextRegionId++
		}
	}
	return regions
}

func Part1(filename string) int {
	regions := scan(readFile(filename))
	totalPrice := 0
	for _, r := range regions {
		perimeter := 0
		for _, p := range r {
			perimeter += p.outerSides
		}
		totalPrice += perimeter * len(r)
	}
	return totalPrice
}

func Part2(filename string) int {
	return 0
}
