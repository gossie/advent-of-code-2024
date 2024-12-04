package day3

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	return scanner.Text()
}

func calculate(s string) int {
	res := 0
	for _, m := range regexp.MustCompile(`mul\(-?\d+,-?\d+\)`).FindAllString(s, -1) {
		left, _ := strconv.Atoi(m[4:strings.Index(m, ",")])
		right, _ := strconv.Atoi(m[strings.Index(m, ",")+1 : strings.Index(m, ")")])
		res += left * right
	}
	return res
}

func Part1(filename string) int {
	content := readFile(filename)
	return calculate(content)
}

func Part2(filename string) int {
	res := 0
	content := readFile(filename)

	for {
		nextDont := strings.Index(content, "don't()")
		if nextDont == -1 {
			res += calculate(content)
			break
		}

		res += calculate(content[:nextDont])
		content = content[nextDont+7:]
		nextDo := strings.Index(content, "do()")
		if nextDo == -1 {
			break
		}
		content = content[nextDo+4:]
	}

	return res
}
