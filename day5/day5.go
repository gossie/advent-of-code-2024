package day5

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type dependencies map[int]map[int]bool

func (r dependencies) addDependency(before, after int) {
	if _, present := r[before]; !present {
		r[before] = make(map[int]bool)
	}
	r[before][after] = true
}

func (r dependencies) dependentOf(before, after int) bool {
	if dependentNodes, present := r[before]; present {
		if _, present := dependentNodes[after]; present {
			return true
		}
	}
	return false
}

type update []int

func (u update) valid(rules dependencies) bool {
	for i := len(u) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if rules.dependentOf(u[i], u[j]) {
				return false
			}
		}
	}
	return true
}

func (u update) sort(rules dependencies) {
	slices.SortFunc(u, func(a, b int) int {
		switch {
		case rules.dependentOf(a, b):
			return -1
		case rules.dependentOf(b, a):
			return 1
		default:
			return 0
		}
	})
}

func (u update) middle() int {
	return u[len(u)/2]
}

func readFile(filename string) (dependencies, []update) {
	file, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer file.Close()

	rules := make(dependencies)
	updates := make([]update, 0)

	scanRules := true
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanRules = false
			continue
		}

		if scanRules {
			pageNumbers := strings.Split(line, "|")
			before, _ := strconv.Atoi(pageNumbers[0])
			after, _ := strconv.Atoi(pageNumbers[1])
			rules.addDependency(before, after)
		} else {
			pageNumbers := strings.Split(line, ",")
			update := make([]int, 0, len(pageNumbers))
			for _, pStr := range pageNumbers {
				page, _ := strconv.Atoi(pStr)
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func Part1(filename string) int {
	rules, updates := readFile(filename)
	sum := 0

	for _, u := range updates {
		if u.valid(rules) {
			sum += u.middle()
		}
	}

	return sum
}

func Part2(filename string) int {
	rules, updates := readFile(filename)
	sum := 0

	for _, u := range updates {
		if !u.valid(rules) {
			u.sort(rules)
			sum += u.middle()
		}
	}

	return sum
}
