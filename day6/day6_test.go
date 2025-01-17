package day6_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day6"
)

func TestPart1(t *testing.T) {
	part1 := day6.Part1("day6_test.txt")
	if part1 != 41 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day6.Part2("day6_test.txt")
	if part2 != 6 {
		t.Fatalf("part2 = %v", part2)
	}
}
