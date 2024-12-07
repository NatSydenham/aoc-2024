package day1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
)

func getValues() ([]int, []int) {
	lines := file.Readlines("./data/day1.txt")

	var first []int
	var second []int
	for _, line := range lines {
		split := strings.Split(line, "   ")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		first = append(first, a)
		second = append(second, b)
	}

	return first, second
}

func ExecutePart1() {
	start := time.Now()
	first, second := getValues()
	sort.Ints(first)
	sort.Ints(second)

	sum := 0

	for i := 0; i < len(first); i++ {
		max := int(math.Max(float64(first[i]), float64(second[i])))
		min := int(math.Min(float64(first[i]), float64(second[i])))
		distance := max - min
		sum += distance
	}
	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	first, second := getValues()
	count := make(map[int]int)

	for i := 0; i < len(second); i++ {
		_, exists := count[second[i]]
		if exists {
			count[second[i]]++
		} else {
			count[second[i]] = 1
		}
	}

	sum := 0

	for i := 0; i < len(first); i++ {
		val, exists := count[first[i]]
		if exists {
			sum += val * first[i]
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}
