package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func validBigger(i, j int) bool {
	return j > i && j-i <= 3
}

func validSmaller(i, j int) bool {
	return i > j && i-j <= 3
}

type report []int

func (r report) errorCount(left, right int, checker func(i, j int) bool) int {
	if left < 0 || right >= len(r) {
		return 0
	}

	if checker(r[left], r[right]) {
		return r.errorCount(right, right+1, checker)
	}

	if left == 0 && right == 1 {
		return 1 + int(math.Min(float64(r.errorCount(1, 2, checker)), float64(r.errorCount(0, 2, checker))))
	}
	return 1 + int(math.Min(float64(r.errorCount(left-1, right, checker)), float64(r.errorCount(left, right+1, checker))))
}

func (r report) safe(allowedErrors int) bool {
	errorsBigger, errorsSmaller := r.errorCount(0, 1, validBigger), r.errorCount(0, 1, validSmaller)
	return errorsBigger <= allowedErrors || errorsSmaller <= allowedErrors
}

func readFile(filename string) []report {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	reports := make([]report, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r := make(report, 0, 5)

		line := scanner.Text()
		levels := strings.Split(line, " ")
		for i := range len(levels) {
			level, _ := strconv.Atoi(levels[i])
			r = append(r, level)
		}
		reports = append(reports, r)
	}

	return reports
}

func Part1(filename string) int {
	reports := readFile(filename)

	safeCount := 0
	for i := range reports {
		if reports[i].safe(0) {
			safeCount++
		}
	}

	return safeCount
}

func Part2(filename string) int {
	reports := readFile(filename)

	safeCount := 0
	for i := range reports {
		if reports[i].safe(1) {
			safeCount++
		}
	}

	return safeCount
}
