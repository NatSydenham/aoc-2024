package day4

import (
	"fmt"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/coordinate"
	"github.com/NatSydenham/aoc-2024/lib/file"
)

func check(input []string, char byte, line int, pos int) bool {
	if line < 0 || line > len(input)-1 || pos > len(input[line])-1 || pos < 0 {
		return false
	}

	return input[line][pos] == char
}

func checkXmas(input []string, line int, pos int, lineIncrement int, posIncrement int) bool {
	return check(input, 'M', line+lineIncrement, pos+posIncrement) &&
		check(input, 'A', line+(2*lineIncrement), pos+(2*posIncrement)) &&
		check(input, 'S', line+(3*lineIncrement), pos+(3*posIncrement))
}

// Checks for M * S or S * M on diagonals
func checkMas(input []string, line int, pos int, lineIncrement int, posIncrement int) bool {
	return check(input, 'M', line, pos) && check(input, 'S', line+(2*lineIncrement), pos+(2*posIncrement)) ||
		check(input, 'S', line, pos) && check(input, 'M', line+(2*lineIncrement), pos+(2*posIncrement))
}

func checkXdash(input []string, line int, pos int) bool {
	return checkMas(input, line-1, pos-1, 1, 1) && checkMas(input, line+1, pos-1, -1, 1)
}

func ExecutePart1() {
	start := time.Now()
	input := file.Readlines("./data/day4.txt")
	sum := 0

	allDirections := [8]coordinate.Coordinate{{X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1}, {X: 0, Y: -1}, {X: -1, Y: -1}}

	for lineno, line := range input {
		for charidx, char := range line {
			if char == 'X' {
				for _, coords := range allDirections {
					if checkXmas(input, lineno, charidx, coords.X, coords.Y) {
						sum++
					}
				}
			}
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	input := file.Readlines("./data/day4.txt")
	sum := 0

	for lineno, line := range input {
		for charidx, char := range line {
			if char == 'A' && checkXdash(input, lineno, charidx) {
				sum++
			}
		}
	}

	fmt.Println(sum, "|", time.Since(start))
}
