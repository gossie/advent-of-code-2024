package day9_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day9"
)

func TestPart1(t *testing.T) {
	part1 := day9.Part1("day9_test.txt")
	if part1 != 1928 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day9.Part2("day9_test.txt")
	if part2 != 2858 {
		t.Fatalf("part2 = %v", part2)
	}
}
