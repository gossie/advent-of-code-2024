package day7

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func add(a, b int64) int64 {
	return a + b
}

func mul(a, b int64) int64 {
	return a * b
}

func concat(a, b int64) int64 {
	return a*padding(b) + b
}

func padding(n int64) int64 {
	if n == 0 {
		return 10
	}

	var p int64 = 1
	for p <= n {
		p *= 10
	}
	return p
}

type equation struct {
	result  int64
	numbers []int64
}

func (e equation) solvable() bool {
	current := []int64{e.numbers[0]}
	for i := 1; i < len(e.numbers); i++ {
		newCurrent := make([]int64, 0, len(current)*2)

		for _, c := range current {
			newCurrent = append(newCurrent, add(c, e.numbers[i]))
			newCurrent = append(newCurrent, mul(c, e.numbers[i]))
		}

		current = newCurrent
	}
	return slices.Contains(current, e.result)
}

func (e equation) solvableWithCombining() bool {
	current := []int64{e.numbers[0]}
	for i := 1; i < len(e.numbers); i++ {
		newCurrent := make([]int64, 0, len(current)*3)

		for _, c := range current {
			newCurrent = append(newCurrent, add(c, e.numbers[i]))
			newCurrent = append(newCurrent, mul(c, e.numbers[i]))
			newCurrent = append(newCurrent, concat(c, e.numbers[i]))
		}

		current = newCurrent
	}
	return slices.Contains(current, e.result)
}

func readFile(filename string) []equation {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	equations := make([]equation, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		eqStr := strings.Split(line, ": ")

		result, _ := strconv.Atoi(eqStr[0])
		numbers := make([]int64, 0, len(eqStr)-1)
		nStr := strings.Split(eqStr[1], " ")
		for _, str := range nStr {
			n, _ := strconv.Atoi(str)
			numbers = append(numbers, int64(n))
		}

		eq := equation{result: int64(result), numbers: numbers}

		equations = append(equations, eq)
	}

	return equations
}

func Part1(filename string) int64 {
	sum := int64(0)
	for _, e := range readFile(filename) {
		if e.solvable() {
			sum += e.result
		}
	}

	return sum
}

func Part2(filename string) int64 {
	sum := int64(0)
	for _, e := range readFile(filename) {
		if e.solvableWithCombining() {
			sum += e.result
		}
	}

	return sum
}
