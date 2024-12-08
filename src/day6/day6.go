package day6

import (
	"fmt"
	"slices"
	"time"

	"github.com/NatSydenham/aoc-2024/lib/coordinate"
	"github.com/NatSydenham/aoc-2024/lib/file"
)

func getStart(lines []string) coordinate.Coordinate {
	for row, line := range lines {
		for idx, val := range line {
			if val == '^' {
				return coordinate.Coordinate{X: idx, Y: row}
			}
		}
	}

	panic("No guard found")
}

func rotateClockwise(direction coordinate.Coordinate) coordinate.Coordinate {
	return coordinate.Coordinate{X: direction.Y * -1, Y: direction.X}
}

func step(lines []string, pos coordinate.Coordinate, direction coordinate.Coordinate) (bool, coordinate.Coordinate, coordinate.Coordinate) {
	next := coordinate.Coordinate{X: pos.X + direction.X, Y: pos.Y + direction.Y}

	if next.Y < 0 || next.Y >= len(lines) || next.X >= len(lines[next.Y]) || next.X < 0 {
		return false, coordinate.Coordinate{X: next.X, Y: next.Y}, direction
	}

	if lines[next.Y][next.X] == '#' {
		newDir := rotateClockwise(direction)
		return true, pos, newDir
	}

	return true, coordinate.Coordinate{X: next.X, Y: next.Y}, coordinate.Coordinate{X: direction.X, Y: direction.Y}
}

func ExecutePart1() {
	start := time.Now()

	lines := file.Readlines("./data/day6.txt")
	visited := make(map[coordinate.Coordinate]bool)
	pos := getStart(lines)

	dir := coordinate.Coordinate{X: 0, Y: -1}
	isInBounds := true

	for isInBounds {
		_, hasVisited := visited[pos]
		if !hasVisited {
			visited[pos] = true
		}

		isInBounds, pos, dir = step(lines, pos, dir)
	}

	fmt.Println(len(visited), "|", time.Since(start))
}

func ExecutePart2() {
	start := time.Now()
	lines := file.Readlines("./data/day6.txt")
	loops := make([]coordinate.Coordinate, 0)
	checked := make([]coordinate.Coordinate, 0)

	obstacleWalkerPos := getStart(lines)
	startPos := obstacleWalkerPos

	obstacleWalkerDir := coordinate.Coordinate{X: 0, Y: -1}
	obstacleWalkerInBounds := true

	for obstacleWalkerInBounds {

		if slices.Contains(checked, obstacleWalkerPos) {
			obstacleWalkerInBounds, obstacleWalkerPos, obstacleWalkerDir = step(lines, obstacleWalkerPos, obstacleWalkerDir)
			continue
		}

		//add obstacle
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		if obstacleWalkerPos != startPos {
			lineWithObs := newLines[obstacleWalkerPos.Y][:obstacleWalkerPos.X] + string("#") + newLines[obstacleWalkerPos.Y][obstacleWalkerPos.X+1:]
			newLines[obstacleWalkerPos.Y] = lineWithObs
		}

		visited := make(map[coordinate.Coordinate][]coordinate.Coordinate)
		guardPos := startPos
		guardInBounds := true
		guardDir := coordinate.Coordinate{X: 0, Y: -1}

		// check if guard loops with added obstacle
		for guardInBounds {
			directions, hasVisited := visited[guardPos]
			if !hasVisited {
				newArr := make([]coordinate.Coordinate, 0)
				visited[guardPos] = append(newArr, guardDir)
			} else if !slices.Contains(directions, guardDir) {
				visited[guardPos] = append(visited[guardPos], guardDir)
			} else {
				// we have a loop
				if !slices.Contains(loops, obstacleWalkerPos) {
					loops = append(loops, obstacleWalkerPos)
				}
				break
			}
			guardInBounds, guardPos, guardDir = step(newLines, guardPos, guardDir)
		}

		// move the obstacle walker to next position in original guard path.
		checked = append(checked, obstacleWalkerPos)
		obstacleWalkerInBounds, obstacleWalkerPos, obstacleWalkerDir = step(lines, obstacleWalkerPos, obstacleWalkerDir)
	}

	fmt.Println(len(loops), "|", time.Since(start))
}
