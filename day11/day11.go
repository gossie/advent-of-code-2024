package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

type stone struct {
	number    int
	numberStr string
}

func readFile(filename string) []stone {
	theFile, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer theFile.Close()

	stones := make([]stone, 0)

	scanner := bufio.NewScanner(theFile)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	for _, n := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(n)
		stones = append(stones, stone{number: number, numberStr: n})
	}

	return stones
}

func calc(s *stone, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := fmt.Sprintf("%v-%v", s.number, blinks)
	if cachedValue, found := cache[key]; found {
		return cachedValue
	}

	res := 0

	switch {
	case s.number == 0:
		res = calc(&stone{number: 1, numberStr: "1"}, blinks-1)
	case len(s.numberStr)%2 == 0:
		half := len(s.numberStr) / 2
		n1, _ := strconv.Atoi(s.numberStr[:half])
		n2, _ := strconv.Atoi(s.numberStr[half:])
		res = calc(&stone{number: n1, numberStr: strconv.Itoa(n1)}, blinks-1) + calc(&stone{number: n2, numberStr: strconv.Itoa(n2)}, blinks-1)
	default:
		newNumber := s.number * 2024
		res = calc(&stone{number: newNumber, numberStr: strconv.Itoa(newNumber)}, blinks-1)
	}
	cache[key] = res

	return res
}

func Part1(filename string) int {
	res := 0
	for _, s := range readFile(filename) {
		res += calc(&s, 25)
	}
	return res
}

func Part2(filename string) int {
	res := 0
	for _, s := range readFile(filename) {
		res += calc(&s, 75)
	}
	return res
}
