package day8

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/coordinate"
	"github.com/NatSydenham/aoc-2024/lib/file"
)

type Pair struct {
	a int
	b int
}

func getAntinodes(a coordinate.Coordinate, b coordinate.Coordinate) (coordinate.Coordinate, coordinate.Coordinate) {
	vector := coordinate.Coordinate{X: b.X - a.X, Y: b.Y - a.Y}

	antinode1 := coordinate.Coordinate{X: b.X + vector.X, Y: b.Y + vector.Y}
	antinode2 := coordinate.Coordinate{X: a.X - vector.X, Y: a.Y - vector.Y}

	return antinode1, antinode2
}

func getAntinodesInDirection(next coordinate.Coordinate, from coordinate.Coordinate, antinodes []coordinate.Coordinate, xLength int, yLength int) []coordinate.Coordinate {
	for {
		antinode := getNextAntinodeInDirection(next, from)
		antinodeIsValid := isValidPosition(antinode, yLength, xLength)

		if antinodeIsValid && !slices.Contains(antinodes, antinode) {
			antinodes = append(antinodes, antinode)
		}

		if !antinodeIsValid {
			break
		}

		from = next
		next = antinode
	}

	return antinodes
}

func getNextAntinodeInDirection(next coordinate.Coordinate, from coordinate.Coordinate) coordinate.Coordinate {
	vector := coordinate.Coordinate{X: from.X - next.X, Y: from.Y - next.Y}
	return coordinate.Coordinate{X: next.X - vector.X, Y: next.Y - vector.Y}
}

func isValidPosition(antinode coordinate.Coordinate, linecount int, linelength int) bool {
	return antinode.X >= 0 && antinode.X < linelength && antinode.Y >= 0 && antinode.Y < linecount
}

func ExecutePart1() {
	start := time.Now()
	lines := file.Readlines("./data/day8.txt")

	symbols := make(map[rune][]coordinate.Coordinate)

	for lineNo, line := range lines {
		for charNo, char := range line {
			if char != '.' {
				arr, exists := symbols[char]
				if exists {
					newArr := append(arr, coordinate.Coordinate{X: charNo, Y: lineNo})
					symbols[char] = newArr
				} else {
					newArr := []coordinate.Coordinate{{X: charNo, Y: lineNo}}
					symbols[char] = newArr
				}
			}
		}
	}

	antinodes := []coordinate.Coordinate{}

	for _, nodes := range symbols {
		if len(nodes) == 1 {
			continue
		}

		pairsDone := []Pair{}

		for i1, node1 := range nodes {
			for i2, node2 := range nodes {

				eval := Pair{a: int(math.Min(float64(i1), float64(i2))), b: int(math.Max(float64(i1), float64(i2)))}

				if i1 == i2 || slices.Contains(pairsDone, eval) {
					continue
				}

				pairsDone = append(pairsDone, eval)
				antinode1, antinode2 := getAntinodes(node1, node2)

				if isValidPosition(antinode1, len(lines), len(lines[0])) && !slices.Contains(antinodes, antinode1) {
					antinodes = append(antinodes, antinode1)
				}

				if isValidPosition(antinode2, len(lines), len(lines[0])) && !slices.Contains(antinodes, antinode2) {
					antinodes = append(antinodes, antinode2)
				}
			}
		}
	}

	fmt.Println(len(antinodes), "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	lines := file.Readlines("./data/day8.txt")

	symbols := make(map[rune][]coordinate.Coordinate)

	for lineNo, line := range lines {
		for charNo, char := range line {
			if char != '.' {
				arr, exists := symbols[char]
				if exists {
					newArr := append(arr, coordinate.Coordinate{X: charNo, Y: lineNo})
					symbols[char] = newArr
				} else {
					newArr := []coordinate.Coordinate{{X: charNo, Y: lineNo}}
					symbols[char] = newArr
				}
			}
		}
	}

	antinodes := []coordinate.Coordinate{}

	for _, nodes := range symbols {
		if len(nodes) == 1 {
			continue
		}

		pairsDone := []Pair{}

		for i1, node1 := range nodes {
			for i2, node2 := range nodes {

				if !slices.Contains(antinodes, node1) {
					antinodes = append(antinodes, node1)
				}

				if !slices.Contains(antinodes, node2) {
					antinodes = append(antinodes, node2)
				}

				eval := Pair{a: int(math.Min(float64(i1), float64(i2))), b: int(math.Max(float64(i1), float64(i2)))}

				if i1 == i2 || slices.Contains(pairsDone, eval) {
					continue
				}

				pairsDone = append(pairsDone, eval)

				antinodes = getAntinodesInDirection(node1, node2, antinodes, len(lines[0]), len(lines))
				antinodes = getAntinodesInDirection(node2, node1, antinodes, len(lines[0]), len(lines))
			}
		}
	}

	fmt.Println(len(antinodes), "|", time.Since(start))

}
