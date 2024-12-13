package day2

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/file"
	"github.com/NatSydenham/aoc-2024/lib/slice"
)

func isOutOfRange(prev int, curr int) bool {
	return prev == curr || prev-curr < -3 || prev-curr > 3
}

func isOutOfOrder(prev int, curr int, isIncreasing bool) bool {
	return (isIncreasing && prev > curr) || (!isIncreasing && prev < curr)
}

func isSafe(line []int) bool {
	isIncreasing := false

	for i, num := range line {
		if i == 0 {
			isIncreasing = line[i] < line[i+1]
			continue
		}
		if isOutOfRange(line[i-1], num) || isOutOfOrder(line[i-1], num, isIncreasing) {
			return false
		}
	}

	return true
}

func ExecutePart1() {
	start := time.Now()
	lines := file.Readlines("./data/day2.txt")

	sum := 0

	for _, line := range lines {
		strLine := strings.Split(line, " ")
		val := make([]int, len(strLine))

		for i, str := range strLine {
			val[i], _ = strconv.Atoi(str)
		}

		safe := isSafe(val)
		if safe {
			sum++
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	lines := file.Readlines("./data/day2.txt")

	sum := 0
	for _, line := range lines {

		strLine := strings.Split(line, " ")
		val := make([]int, len(strLine))

		for i, str := range strLine {
			val[i], _ = strconv.Atoi(str)
		}

		safe := isSafe(val)

		if safe {
			sum++
		} else {
			for i := range val {
				updated := slice.RemoveAt(val, i)
				safeAfterAdjust := isSafe(updated)
				if safeAfterAdjust {
					sum++
					break
				}
			}
		}
	}

	fmt.Println(sum, "|", time.Since(start))

}
