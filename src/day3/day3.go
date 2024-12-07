package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
)

func calcSum(input []string) int {
	sum := 0

	for _, val := range input {
		parts := strings.Split(val, ",")
		a, _ := strconv.Atoi(parts[0][4:len(parts[0])])
		b, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ")"))

		sum += (a * b)
	}

	return sum
}

func ExecutePart1() {
	start := time.Now()
	input := strings.Join(file.
		Readlines("./data/day3.txt"), "")
	r := regexp.
		MustCompile("mul\\(\\d+,\\d+\\)")
	matches := r.FindAllString(input, -1)
	sum := calcSum(matches)
	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	input := strings.Join(file.Readlines("./data/day3.txt"), "")
	validInstructions := regexp.MustCompile("don't\\(\\).*").
		ReplaceAllString(
			regexp.MustCompile("don't\\(\\).*?do\\(\\)").
				ReplaceAllString(input, ""), "")
	r := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := r.FindAllString(validInstructions, -1)
	sum := calcSum(matches)
	fmt.Println(sum, "|", time.Since(start))
}
