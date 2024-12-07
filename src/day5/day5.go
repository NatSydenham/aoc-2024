package day5

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
	"github.com/NatSydenham/aoc-2024/lib/slice"
)

func getInput() ([][]int, [][]int) {
	rulesRaw, inputRaw := file.Readlines("./data/day5rules.txt"), file.Readlines("./data/day5input.txt")

	rules := slice.Map(rulesRaw, func(rule string) []int {
		return slice.Map(strings.Split(rule, "|"), func(r string) int {
			conv, _ := strconv.Atoi(r)
			return conv
		})
	})

	input := slice.Map(inputRaw, func(line string) []int {
		return slice.Map(strings.Split(line, ","), func(l string) int {
			conv, _ := strconv.Atoi(l)
			return conv
		})
	})

	return rules, input
}

func check(rule []int, line []int) bool {
	posA, posB := slice.IndexOf(rule[0], line), slice.IndexOf(rule[1], line)
	return (posA == -1 || posB == -1) || posA < posB

}

func ExecutePart1() {
	start := time.Now()
	rules, input := getInput()
	sum := 0

	for _, line := range input {
		passesAllRules := true

		for _, rule := range rules {
			if !check(rule, line) {
				passesAllRules = false
				continue
			}
		}

		if passesAllRules {
			sum += line[len(line)/2]
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	rules, input := getInput()
	sum := 0

	swaps := 0

	for _, line := range input {
		passesAllRules := true
		shouldCheck := true

		for shouldCheck {
			shouldCheck = false
			for _, rule := range rules {
				if !check(rule, line) {
					passesAllRules = false
					shouldCheck = true
					aPos, bPos := slice.IndexOf(rule[0], line), slice.IndexOf(rule[1], line)
					line[aPos], line[bPos] = line[bPos], line[aPos]
					swaps++
				}
			}
		}

		if !passesAllRules {
			sum += line[len(line)/2]
		}
	}

	fmt.Println(sum, "|", time.Since(start), "|", swaps, "swaps")
}
