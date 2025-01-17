package day2_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day2"
)

func TestPart1(t *testing.T) {
	part1 := day2.Part1("day2_test.txt")
	if part1 != 2 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day2.Part2("day2_test.txt")
	if part2 != 4 {
		t.Fatalf("part2 = %v", part2)
	}
}
