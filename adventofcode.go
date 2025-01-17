package main

import (
	"fmt"
	"time"

	"github.com/gossie/adventofcode2024/day1"
	"github.com/gossie/adventofcode2024/day2"
	"github.com/gossie/adventofcode2024/day3"
	"github.com/gossie/adventofcode2024/day4"
	"github.com/gossie/adventofcode2024/day5"
	"github.com/gossie/adventofcode2024/day6"
	"github.com/gossie/adventofcode2024/day7"
	"github.com/gossie/adventofcode2024/day8"
	"github.com/gossie/adventofcode2024/day9"
    "github.com/gossie/adventofcode2024/day10"
    "github.com/gossie/adventofcode2024/day11"
    "github.com/gossie/adventofcode2024/day12"
)

func main() {

	fmt.Println("\nPerforming tasks of day1")
	startday1Part1 := time.Now().UnixMilli()
	day1Part1 := day1.Part1("day1/day1.txt")
	fmt.Println("day1, task 1: ", day1Part1, ", took", (time.Now().UnixMilli() - startday1Part1), "ms")
	startday1Part2 := time.Now().UnixMilli()
	day1Part2 := day1.Part2("day1/day1.txt")
	fmt.Println("day1, task 2: ", day1Part2, ", took", (time.Now().UnixMilli() - startday1Part2), "ms")

	fmt.Println("\nPerforming tasks of day2")
	startday2Part1 := time.Now().UnixMilli()
	day2Part1 := day2.Part1("day2/day2.txt")
	fmt.Println("day2, task 1: ", day2Part1, ", took", (time.Now().UnixMilli() - startday2Part1), "ms")
	startday2Part2 := time.Now().UnixMilli()
	day2Part2 := day2.Part2("day2/day2.txt")
	fmt.Println("day2, task 2: ", day2Part2, ", took", (time.Now().UnixMilli() - startday2Part2), "ms")

	fmt.Println("\nPerforming tasks of day3")
	startday3Part1 := time.Now().UnixMilli()
	day3Part1 := day3.Part1("day3/day3.txt")
	fmt.Println("day3, task 1: ", day3Part1, ", took", (time.Now().UnixMilli() - startday3Part1), "ms")
	startday3Part2 := time.Now().UnixMilli()
	day3Part2 := day3.Part2("day3/day3.txt")
	fmt.Println("day3, task 2: ", day3Part2, ", took", (time.Now().UnixMilli() - startday3Part2), "ms")

	fmt.Println("\nPerforming tasks of day4")
	startday4Part1 := time.Now().UnixMilli()
	day4Part1 := day4.Part1("day4/day4.txt")
	fmt.Println("day4, task 1: ", day4Part1, ", took", (time.Now().UnixMilli() - startday4Part1), "ms")
	startday4Part2 := time.Now().UnixMilli()
	day4Part2 := day4.Part2("day4/day4.txt")
	fmt.Println("day4, task 2: ", day4Part2, ", took", (time.Now().UnixMilli() - startday4Part2), "ms")

	fmt.Println("\nPerforming tasks of day5")
	startday5Part1 := time.Now().UnixMilli()
	day5Part1 := day5.Part1("day5/day5.txt")
	fmt.Println("day5, task 1: ", day5Part1, ", took", (time.Now().UnixMilli() - startday5Part1), "ms")
	startday5Part2 := time.Now().UnixMilli()
	day5Part2 := day5.Part2("day5/day5.txt")
	fmt.Println("day5, task 2: ", day5Part2, ", took", (time.Now().UnixMilli() - startday5Part2), "ms")

	fmt.Println("\nPerforming tasks of day6")
	startday6Part1 := time.Now().UnixMilli()
	day6Part1 := day6.Part1("day6/day6.txt")
	fmt.Println("day6, task 1: ", day6Part1, ", took", (time.Now().UnixMilli() - startday6Part1), "ms")
	startday6Part2 := time.Now().UnixMilli()
	day6Part2 := day6.Part2("day6/day6.txt")
	fmt.Println("day6, task 2: ", day6Part2, ", took", (time.Now().UnixMilli() - startday6Part2), "ms")

	fmt.Println("\nPerforming tasks of day7")
	startday7Part1 := time.Now().UnixMilli()
	day7Part1 := day7.Part1("day7/day7.txt")
	fmt.Println("day7, task 1: ", day7Part1, ", took", (time.Now().UnixMilli() - startday7Part1), "ms")
	startday7Part2 := time.Now().UnixMilli()
	day7Part2 := day7.Part2("day7/day7.txt")
	fmt.Println("day7, task 2: ", day7Part2, ", took", (time.Now().UnixMilli() - startday7Part2), "ms")

	fmt.Println("\nPerforming tasks of day8")
	startday8Part1 := time.Now().UnixMilli()
	day8Part1 := day8.Part1("day8/day8.txt")
	fmt.Println("day8, task 1: ", day8Part1, ", took", (time.Now().UnixMilli() - startday8Part1), "ms")
	startday8Part2 := time.Now().UnixMilli()
	day8Part2 := day8.Part2("day8/day8.txt")
	fmt.Println("day8, task 2: ", day8Part2, ", took", (time.Now().UnixMilli() - startday8Part2), "ms")

	fmt.Println("\nPerforming tasks of day9")
	startday9Part1 := time.Now().UnixMilli()
	day9Part1 := day9.Part1("day9/day9.txt")
	fmt.Println("day9, task 1: ", day9Part1, ", took", (time.Now().UnixMilli() - startday9Part1), "ms")
	startday9Part2 := time.Now().UnixMilli()
	day9Part2 := day9.Part2("day9/day9.txt")
	fmt.Println("day9, task 2: ", day9Part2, ", took", (time.Now().UnixMilli() - startday9Part2), "ms")

    fmt.Println("\nPerforming tasks of day10")
    startday10Part1 := time.Now().UnixMilli()
    day10Part1 := day10.Part1("day10/day10.txt")
    fmt.Println("day10, task 1: ", day10Part1, ", took", (time.Now().UnixMilli() - startday10Part1), "ms")
    startday10Part2 := time.Now().UnixMilli()
    day10Part2 := day10.Part2("day10/day10.txt")
    fmt.Println("day10, task 2: ", day10Part2, ", took", (time.Now().UnixMilli() - startday10Part2), "ms")

    fmt.Println("\nPerforming tasks of day11")
    startday11Part1 := time.Now().UnixMilli()
    day11Part1 := day11.Part1("day11/day11.txt")
    fmt.Println("day11, task 1: ", day11Part1, ", took", (time.Now().UnixMilli() - startday11Part1), "ms")
    startday11Part2 := time.Now().UnixMilli()
    day11Part2 := day11.Part2("day11/day11.txt")
    fmt.Println("day11, task 2: ", day11Part2, ", took", (time.Now().UnixMilli() - startday11Part2), "ms")

    fmt.Println("\nPerforming tasks of day12")
    startday12Part1 := time.Now().UnixMilli()
    day12Part1 := day12.Part1("day12/day12.txt")
    fmt.Println("day12, task 1: ", day12Part1, ", took", (time.Now().UnixMilli() - startday12Part1), "ms")
    startday12Part2 := time.Now().UnixMilli()
    day12Part2 := day12.Part2("day12/day12.txt")
    fmt.Println("day12, task 2: ", day12Part2, ", took", (time.Now().UnixMilli() - startday12Part2), "ms")
}
