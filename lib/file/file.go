package file

import (
	"bufio"
	"os"
)

func Readlines(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
