package day1

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		l, _ := strconv.Atoi(numbers[0])
		r, _ := strconv.Atoi(numbers[1])
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func Part1(filename string) int {
	left, right := readFile(filename)
	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	return sum
}

func Part2(filename string) int {
	left, right := readFile(filename)
	rightCounted := make(map[int]int)

	for _, v := range right {
		rightCounted[v]++
	}

	sum := 0

	for _, v := range left {
		sum += v * rightCounted[v]
	}

	return sum
}
