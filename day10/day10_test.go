package day10_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day10"
)

func TestPart1(t *testing.T) {
	part1 := day10.Part1("day10_test.txt")
	if part1 != 36 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day10.Part2("day10_test.txt")
	if part2 != 81 {
		t.Fatalf("part2 = %v", part2)
	}
}
