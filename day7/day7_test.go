package day7_test

import (
	"testing"

	"github.com/gossie/adventofcode2024/day7"
)

func TestPart1(t *testing.T) {
	part1 := day7.Part1("day7_test.txt")
	if part1 != 3749 {
		t.Fatalf("part1 = %v", part1)
	}
}

func TestPart2(t *testing.T) {
	part2 := day7.Part2("day7_test.txt")
	if part2 != 11387 {
		t.Fatalf("part2 = %v", part2)
	}
}
