package day3_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day3"
)

func TestPart1(t *testing.T) {
	part1 := day3.Part1("day3_test.txt")
	if part1 != 161 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day3.Part2("day3_test.txt")
	if part2 != 48 {
		t.Fatalf("part2 = %v", part2)
	}
}
