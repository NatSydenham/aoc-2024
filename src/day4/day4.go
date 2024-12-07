package day4

import (
	"fmt"

	"github.com/NatSydenham/aoc-2024/lib/file"
)

type Coordinate struct {
	x, y int
}

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

func ExecutePart1() {
	input := file.Readlines("./data/day4.txt")
	sum := 0

	for lineno, line := range input {
		for charidx, char := range line {
			if char == 'X' {
				for _, coords := range [8]Coordinate{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}} {

					if checkXmas(input, lineno, charidx, coords.x, coords.y) {
						fmt.Println(lineno, charidx, coords.x, coords.y)
						sum++
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
