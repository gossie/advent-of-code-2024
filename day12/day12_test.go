package day12_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day12"
)

func TestPart1(t *testing.T) {
	part1 := day12.Part1("day12_test.txt")
	if part1 != 1930 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day12.Part2("day12_test.txt")
	if part2 != 0 {
		t.Fatalf("part2 = %v", part2)
	}
}
